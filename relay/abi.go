package relay

import (
	"context"
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
	eventABI    abi.ABI
	factoryMsg  evmTypes.ExecuteFactoryMsg
	cost        int64
	errUnpack   error
	errGenerate error
)

func init() {
	var err error
	eventABI, err = abi.JSON(strings.NewReader(ABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}
}

func generatePostedFileMsg(w *wallet.Wallet, q *uploader.Queue, chainID uint64, jackalContract string, event PostedFile) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress := event.Sender.String()

	merkleRoot, err := hex.DecodeString(event.Merkle)
	if err != nil {
		log.Printf("Failed to decode merkle: %v", err)
		return
	}

	abci, err := w.Client.RPCClient.ABCIInfo(context.Background())
	if err != nil {
		log.Printf("Failed to query ABCI: %v", err)
		return
	}

	log.Printf("Relaying for %s\n", event.Sender.String())

	merkleBase64 := base64.StdEncoding.EncodeToString(merkleRoot)
	var maxProofs int64 = 3
	fileSize := int64(event.Size)

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

	storageMsg := evmTypes.ExecuteMsg{
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

	factoryMsg = evmTypes.ExecuteFactoryMsg{
		CallBindings: &evmTypes.ExecuteMsgCallBindings{
			EvmAddress: &evmAddress,
			Msg:        &storageMsg,
		},
	}

	cost = q.GetCost(fileSize*maxProofs, int64(event.Expires)*24)
	cost = int64(float64(cost) * 1.2)
	return
}

func generateBoughtStorageMsg(w *wallet.Wallet, q *uploader.Queue, chainID uint64, jackalContract string, event BoughtStorage) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress := event.From.String()
	log.Printf("Relaying for %s\n", event.From.String())

	storageMsg := evmTypes.ExecuteMsg{
		BuyStorage: &evmTypes.ExecuteMsgBuyStorage{
			ForAddress:   event.ForAddress,
			DurationDays: int64(event.DurationDays),
			Bytes:        int64(event.SizeBytes),
			PaymentDenom: "ujkl",
			Referral:     event.Referral,
		},
	}

	factoryMsg = evmTypes.ExecuteFactoryMsg{
		CallBindings: &evmTypes.ExecuteMsgCallBindings{
			EvmAddress: &evmAddress,
			Msg:        &storageMsg,
		},
	}

	cost = q.GetCost(int64(event.SizeBytes), int64(event.DurationDays)*24) // double check cost calculation
	cost = int64(float64(cost) * 1.2)
	return
}

func generateDeletedFileMsg(w *wallet.Wallet, q *uploader.Queue, chainID uint64, jackalContract string, event DeletedFile) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress := event.From.String()
	log.Printf("Relaying for %s\n", event.From.String())

	merkleRoot, err := hex.DecodeString(event.Merkle)
	if err != nil {
		log.Printf("Failed to decode merkle: %v", err)
		return
	}
	merkleBase64 := base64.StdEncoding.EncodeToString(merkleRoot)

	storageMsg := evmTypes.ExecuteMsg{
		DeleteFile: &evmTypes.ExecuteMsgDeleteFile{
			Merkle: merkleBase64,
			Start:  int64(event.Start),
		},
	}

	factoryMsg = evmTypes.ExecuteFactoryMsg{
		CallBindings: &evmTypes.ExecuteMsgCallBindings{
			EvmAddress: &evmAddress,
			Msg:        &storageMsg,
		},
	}

	cost = q.GetCost(0, 0) // minimum nonzero cost
	cost = int64(float64(cost) * 1.2)
	return
}

func handleLog(vLog *types.Log, w *wallet.Wallet, q *uploader.Queue, chainID uint64, jackalContract string) {
	/*
		e, err := eventABI.Unpack("PostedFile", vLog.Data)
		if err != nil {
			log.Fatalf("Failed to unpack log data normally: %v", err)
			return
		}
		fmt.Println(len(e))
	*/

	// can't if-elif-else or case-switch because we need logging
	eventPostedFile := PostedFile{}
	eventBoughtStorage := BoughtStorage{}
	eventDeletedFile := DeletedFile{}

	if errUnpack = eventABI.UnpackIntoInterface(&eventPostedFile, "PostedFile", vLog.Data); errUnpack == nil {
		if errGenerate = generatePostedFileMsg(w, q, chainID, jackalContract, eventPostedFile); errGenerate == nil {
			goto execute
		}
	}
	log.Printf("Failed to unpack log data into PostedFile: %v %v", errUnpack, errGenerate)

	if errUnpack = eventABI.UnpackIntoInterface(&eventBoughtStorage, "BoughtStorage", vLog.Data); errUnpack == nil {
		if errGenerate = generateBoughtStorageMsg(w, q, chainID, jackalContract, eventBoughtStorage); errGenerate == nil {
			goto execute
		}
	}
	log.Printf("Failed to unpack log data into BoughtStorage: %v  %v", errUnpack, errGenerate)

	if errUnpack = eventABI.UnpackIntoInterface(&eventDeletedFile, "DeletedFile", vLog.Data); errUnpack == nil {
		if errGenerate = generateDeletedFileMsg(w, q, chainID, jackalContract, eventDeletedFile); errGenerate == nil {
			goto execute
		}
	}
	log.Printf("Failed to unpack log data into DeletedFile: %v  %v", errUnpack, errGenerate)

	log.Fatalf("Failed to unpack log data into all event types: %v", errUnpack)
	return

execute:
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
		log.Fatalf("something went wrong, response is empty")
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
