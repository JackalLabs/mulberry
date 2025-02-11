package relay

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	_ "embed"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/JackalLabs/mulberry/jackal/uploader"
	evmTypes "github.com/JackalLabs/mulberry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/cosmos-go-wallet/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

//go:embed abi.json
var ABI string

// from `forge inspect Jackal abi`

var (
	eventABI abi.ABI
	errABI   error

	evmAddress string
	cost       int64
	relayedMsg evmTypes.ExecuteMsg
	factoryMsg evmTypes.ExecuteFactoryMsg

	errUnpack   error
	errGenerate error
)

func init() {
	eventABI, errABI = abi.JSON(strings.NewReader(ABI))
	if errABI != nil {
		log.Fatalf("Failed to parse ABI: %v", errABI)
	}
}

func generatePostedFileMsg(w *wallet.Wallet, q *uploader.Queue, chainID uint64, event PostedFile) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	merkleBase64, err := merkleToString(event.Merkle)
	if err != nil {
		return
	}

	abci, err := w.Client.RPCClient.ABCIInfo(context.Background())
	if err != nil {
		log.Printf("Failed to query ABCI: %v", err)
		return
	}

	note := make(map[string]any)
	err = json.Unmarshal([]byte(event.Note), &note)
	if err != nil && event.Note != "" {
		log.Printf("Could not unmarshal: %v", err)
		return
	}

	note["relayed"] = map[string]any{"chain_id": chainRep(chainID), "for": evmAddress}
	newNote, err := json.Marshal(note)
	if err != nil {
		log.Printf("Failed to add memo: %v", err)
		return
	}

	// calculate expires field (event.Expires is the number of days)
	expires := int64(0)
	if event.Expires != 0 {
		expires = abci.Response.LastBlockHeight + ((int64(event.Expires) * 24 * 60 * 60) / 6)
	}

	// fileSize and maxProofs (total storage used)
	fileSize, maxProofs := int64(event.Size), int64(3)

	relayedMsg = evmTypes.ExecuteMsg{
		PostFile: &evmTypes.ExecuteMsgPostFile{
			Merkle:        merkleBase64,
			FileSize:      fileSize,
			ProofInterval: 7200,
			ProofType:     0,
			MaxProofs:     maxProofs,
			Note:          string(newNote),
			Expires:       expires,
		},
	}

	cost = int64(float64(q.GetCost(fileSize*maxProofs, int64(event.Expires)*24)) * 1.2)
	return
}

func generateBoughtStorageMsg(q *uploader.Queue, event BoughtStorage) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		BuyStorage: &evmTypes.ExecuteMsgBuyStorage{
			ForAddress:   event.ForAddress,
			DurationDays: int64(event.DurationDays),
			Bytes:        int64(event.SizeBytes),
			PaymentDenom: "ujkl",
			Referral:     event.Referral,
		},
	}

	cost = int64(float64(q.GetCost(int64(event.SizeBytes), int64(event.DurationDays)*24)) * 1.2)
	return
}

func generateDeletedFileMsg(q *uploader.Queue, event DeletedFile) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	merkleBase64, err := merkleToString(event.Merkle)
	if err != nil {
		return
	}

	relayedMsg = evmTypes.ExecuteMsg{
		DeleteFile: &evmTypes.ExecuteMsgDeleteFile{
			Merkle: merkleBase64,
			Start:  int64(event.Start),
		},
	}

	cost = int64(float64(q.GetCost(0, 0)) * 1.2) // minimum nonzero cost
	return
}

func generateRequestedReportFormMsg(event RequestedReportForm) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	merkleBase64, err := merkleToString(event.Merkle)
	if err != nil {
		return
	}

	relayedMsg = evmTypes.ExecuteMsg{
		RequestReportForm: &evmTypes.ExecuteMsgRequestReportForm{
			Prover: event.Prover,
			Merkle: merkleBase64,
			Owner:  event.Owner,
			Start:  int64(event.Start),
		},
	}
	return
}

func generatePostedKeyMsg(q *uploader.Queue, event PostedKey) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		PostKey: &evmTypes.ExecuteMsgPostKey{
			Key: event.Key,
		},
	}

	cost = int64(float64(q.GetCost(0, 0)) * 1.2) // minimum nonzero cost
	return
}

func generateDeletedFileTreeMsg(q *uploader.Queue, event DeletedFileTree) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	h := sha256.New()
	h.Write([]byte(event.Account))
	accountHash := h.Sum(nil)

	relayedMsg = evmTypes.ExecuteMsg{
		DeleteFileTree: &evmTypes.ExecuteMsgDeleteFileTree{
			HashPath: MerklePath(event.HashPath),
			Account:  fmt.Sprintf("%x", accountHash),
		},
	}

	cost = int64(float64(q.GetCost(0, 0)) * 1.2) // minimum nonzero cost
	return
}

func handleLog(vLog *types.Log, w *wallet.Wallet, q *uploader.Queue, chainID uint64, jackalContract string) {
	// can't if-elif-else or case-switch because we need logging in between
	eventPostedFile := PostedFile{}
	eventBoughtStorage := BoughtStorage{}
	eventDeletedFile := DeletedFile{}
	eventRequestedReportForm := RequestedReportForm{}
	eventDeletedFileTree := DeletedFileTree{}
	eventPostedKey := PostedKey{}

	if errUnpack = eventABI.UnpackIntoInterface(&eventPostedFile, "PostedFile", vLog.Data); errUnpack == nil {
		if errGenerate = generatePostedFileMsg(w, q, chainID, eventPostedFile); errGenerate == nil {
			goto execute
		}
	}
	log.Printf("Failed to unpack log data into PostedFile: %v %v", errUnpack, errGenerate)

	if errUnpack = eventABI.UnpackIntoInterface(&eventBoughtStorage, "BoughtStorage", vLog.Data); errUnpack == nil {
		if errGenerate = generateBoughtStorageMsg(q, eventBoughtStorage); errGenerate == nil {
			goto execute
		}
	}
	log.Printf("Failed to unpack log data into BoughtStorage: %v  %v", errUnpack, errGenerate)

	if errUnpack = eventABI.UnpackIntoInterface(&eventDeletedFile, "DeletedFile", vLog.Data); errUnpack == nil {
		if errGenerate = generateDeletedFileMsg(q, eventDeletedFile); errGenerate == nil {
			goto execute
		}
	}
	log.Printf("Failed to unpack log data into DeletedFile: %v  %v", errUnpack, errGenerate)

	if errUnpack = eventABI.UnpackIntoInterface(&eventRequestedReportForm, "RequestedReportForm", vLog.Data); errUnpack == nil {
		if errGenerate = generateRequestedReportFormMsg(eventRequestedReportForm); errGenerate == nil {
			goto execute
		}
	}
	log.Printf("Failed to unpack log data into RequestedReportForm: %v  %v", errUnpack, errGenerate)

	if errUnpack = eventABI.UnpackIntoInterface(&eventDeletedFileTree, "DeletedFileTree", vLog.Data); errUnpack == nil {
		if errGenerate = generateDeletedFileTreeMsg(q, eventDeletedFileTree); errGenerate == nil {
			goto execute
		}
	}
	log.Printf("Failed to unpack log data into DeletedFileTree: %v  %v", errUnpack, errGenerate)

	if errUnpack = eventABI.UnpackIntoInterface(&eventPostedKey, "PostedKey", vLog.Data); errUnpack == nil {
		if errGenerate = generatePostedKeyMsg(q, eventPostedKey); errGenerate == nil {
			goto execute
		}
	}
	log.Printf("Failed to unpack log data into PostedKey: %v  %v", errUnpack, errGenerate)

	log.Fatal("Failed to unpack log data into all event types")

execute:
	factoryMsg = evmTypes.ExecuteFactoryMsg{
		CallBindings: &evmTypes.ExecuteMsgCallBindings{
			EvmAddress: &evmAddress,
			Msg:        &relayedMsg,
		},
	}

	msg := &wasm.MsgExecuteContract{
		Sender:   w.AccAddress(),
		Contract: jackalContract,
		Msg:      factoryMsg.Encode(),
		Funds:    sdk.NewCoins(sdk.NewInt64Coin("ujkl", cost)),
	}

	log.Printf("execute msg: %v", msg)
	if err := msg.ValidateBasic(); err != nil {
		log.Fatalf("Failed to validate message: %v", err)
		return
	}

	res, err := q.Post(msg)
	if err != nil {
		log.Fatalf("Failed to post message: %v", err)
		return
	}
	if res == nil {
		log.Fatalf("Response is empty")
		return
	}

	log.Println(res.RawLog)
	log.Println(res.TxHash)
}

func chainRep(id uint64) string {
	s := ChainIDS[id]
	if len(s) == 0 {
		return fmt.Sprintf("%d", id)
	}
	return s
}

func merkleToString(merkle string) (string, error) {
	merkleRoot, err := hex.DecodeString(merkle)
	if err != nil {
		log.Printf("Failed to decode merkle: %v", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(merkleRoot), nil
}

func MerklePath(path string) (total string) { // ex: hello/world/path -> ["hello", "world", "path"] -> 3867baa2724c672442e4ba21b6fa532a6380d06a2f8779f11d626bd840d1cdee
	trimPath := strings.TrimSuffix(path, "/")
	chunks := strings.Split(trimPath, "/")

	for _, chunk := range chunks {
		h := sha256.New()
		h.Write([]byte(chunk))
		b := fmt.Sprintf("%x", h.Sum(nil))
		k := fmt.Sprintf("%s%s", total, b)

		h1 := sha256.New()
		h1.Write([]byte(k))
		total = fmt.Sprintf("%x", h1.Sum(nil))
	}
	return
}
