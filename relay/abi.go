package relay

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	_ "embed"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/JackalLabs/mulberry/jackal/uploader"
	evmTypes "github.com/JackalLabs/mulberry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/cosmos-go-wallet/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
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

func generateDeletedFileMsg(event DeletedFile) (err error) {
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

func generatePostedKeyMsg(event PostedKey) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		PostKey: &evmTypes.ExecuteMsgPostKey{
			Key: event.Key,
		},
	}
	return
}

func generateDeletedFileTreeMsg(event DeletedFileTree) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		DeleteFileTree: &evmTypes.ExecuteMsgDeleteFileTree{
			HashPath: event.HashPath,
			Account:  event.Account,
		},
	}
	return
}

func generateProvisionedFiletreeMsg(event ProvisionedFileTree) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		ProvisionFileTree: &evmTypes.ExecuteMsgProvisionFileTree{
			Editors:        event.Editors,
			Viewers:        event.Viewers,
			TrackingNumber: event.TrackingNumber,
		},
	}
	return
}

func generatePostedFileTreeMsg(event PostedFileTree) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		PostFileTree: &evmTypes.ExecuteMsgPostFileTree{
			Account:        event.Account,
			HashParent:     event.HashParent,
			HashChild:      event.HashChild,
			Contents:       event.Contents,
			Viewers:        event.Viewers,
			Editors:        event.Editors,
			TrackingNumber: event.TrackingNumber,
		},
	}
	return
}

func generateAddedViewersMsg(event AddedViewers) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		AddViewers: &evmTypes.ExecuteMsgAddViewers{
			ViewerIds:  event.ViewerIds,
			ViewerKeys: event.ViewerKeys,
			Address:    event.ForAddress,
			FileOwner:  event.FileOwner,
		},
	}
	return
}

func generateRemovedViewersMsg(event RemovedViewers) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		RemoveViewers: &evmTypes.ExecuteMsgRemoveViewers{
			ViewerIds: event.ViewerIds,
			Address:   event.ForAddress,
			FileOwner: event.FileOwner,
		},
	}
	return
}

func generateResetViewersMsg(event ResetViewers) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		ResetViewers: &evmTypes.ExecuteMsgResetViewers{
			Address:   event.ForAddress,
			FileOwner: event.FileOwner,
		},
	}
	return
}

func generateChangedOwnerMsg(event ChangedOwner) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		ChangeOwner: &evmTypes.ExecuteMsgChangeOwner{
			Address:   event.ForAddress,
			FileOwner: event.FileOwner,
			NewOwner:  event.NewOwner,
		},
	}
	return
}

func generateAddedEditorsMsg(event AddedEditors) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		AddEditors: &evmTypes.ExecuteMsgAddEditors{
			EditorIds:  event.EditorIds,
			EditorKeys: event.EditorKeys,
			Address:    event.ForAddress,
			FileOwner:  event.FileOwner,
		},
	}
	return
}

func generateRemovedEditorsMsg(event RemovedEditors) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		RemoveEditors: &evmTypes.ExecuteMsgRemoveEditors{
			EditorIds: event.EditorIds,
			Address:   event.ForAddress,
			FileOwner: event.FileOwner,
		},
	}
	return
}

func generateResetEditorsMsg(event ResetEditors) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		ResetEditors: &evmTypes.ExecuteMsgResetEditors{
			Address:   event.ForAddress,
			FileOwner: event.FileOwner,
		},
	}
	return
}

func generateCreatedNotificationMsg(event CreatedNotification) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		CreateNotification: &evmTypes.ExecuteMsgCreateNotification{
			To:              event.To,
			Contents:        event.Contents,
			PrivateContents: event.PrivateContents,
		},
	}
	return
}

func generateDeletedNotificationMsg(event DeletedNotification) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		DeleteNotification: &evmTypes.ExecuteMsgDeleteNotification{
			From: event.NotificationFrom,
			Time: int64(event.Time),
		},
	}
	return
}

func generateBlockedSendersMsg(event BlockedSenders) (err error) {
	log.Printf("Event details: %+v", event)
	evmAddress = event.From.String()

	relayedMsg = evmTypes.ExecuteMsg{
		BlockedSenders: &evmTypes.ExecuteMsgBlockedSenders{
			ToBlock: event.ToBlock,
		},
	}
	return
}

func handleLog(vLog *types.Log, w *wallet.Wallet, wEth *hdwallet.Wallet, q *uploader.Queue, chainID uint64, RPC string, jackalContract string) {
	// https://goethereumbook.org/event-read/#topics
	eventSig := vLog.Topics[0].Hex()
	switch eventSig {
	case expectedSig("PostedFile(address,string,uint64,string,uint64)"):
		eventPostedFile := PostedFile{}
		errUnpack = eventABI.UnpackIntoInterface(&eventPostedFile, "PostedFile", vLog.Data)
		errGenerate = generatePostedFileMsg(w, q, chainID, eventPostedFile)
	case expectedSig("BoughtStorage(address,string,uint64,uint64,string)"):
		eventBoughtStorage := BoughtStorage{}
		errUnpack = eventABI.UnpackIntoInterface(&eventBoughtStorage, "BoughtStorage", vLog.Data)
		errGenerate = generateBoughtStorageMsg(q, eventBoughtStorage)
	case expectedSig("DeletedFile(address,string,uint64)"):
		eventDeletedFile := DeletedFile{}
		errUnpack = eventABI.UnpackIntoInterface(&eventDeletedFile, "DeletedFile", vLog.Data)
		errGenerate = generateDeletedFileMsg(eventDeletedFile)
	case expectedSig("RequestedReportForm(address,string,string,string,uint64)"):
		eventRequestedReportForm := RequestedReportForm{}
		errUnpack = eventABI.UnpackIntoInterface(&eventRequestedReportForm, "RequestedReportForm", vLog.Data)
		errGenerate = generateRequestedReportFormMsg(eventRequestedReportForm)
	case expectedSig("PostedKey(address,string)"):
		eventPostedKey := PostedKey{}
		errUnpack = eventABI.UnpackIntoInterface(&eventPostedKey, "PostedKey", vLog.Data)
		errGenerate = generatePostedKeyMsg(eventPostedKey)
	case expectedSig("DeletedFileTree(address,string,string)"):
		eventDeletedFileTree := DeletedFileTree{}
		errUnpack = eventABI.UnpackIntoInterface(&eventDeletedFileTree, "DeletedFileTree", vLog.Data)
		errGenerate = generateDeletedFileTreeMsg(eventDeletedFileTree)
	case expectedSig("ProvisionedFileTree(address,string,string,string)"):
		eventProvisionedFileTree := ProvisionedFileTree{}
		errUnpack = eventABI.UnpackIntoInterface(&eventProvisionedFileTree, "ProvisionedFileTree", vLog.Data)
		errGenerate = generateProvisionedFiletreeMsg(eventProvisionedFileTree)
	case expectedSig("PostedFileTree(address,string,string,string,string,string,string,string)"):
		eventPostedFileTree := PostedFileTree{}
		errUnpack = eventABI.UnpackIntoInterface(&eventPostedFileTree, "PostedFileTree", vLog.Data)
		errGenerate = generatePostedFileTreeMsg(eventPostedFileTree)
	case expectedSig("AddedViewers(address,string,string,string,string)"):
		eventAddedViewers := AddedViewers{}
		errUnpack = eventABI.UnpackIntoInterface(&eventAddedViewers, "AddedViewers", vLog.Data)
		errGenerate = generateAddedViewersMsg(eventAddedViewers)
	case expectedSig("RemovedViewers(address,string,string,string)"):
		eventRemovedViewers := RemovedViewers{}
		errUnpack = eventABI.UnpackIntoInterface(&eventRemovedViewers, "RemovedViewers", vLog.Data)
		errGenerate = generateRemovedViewersMsg(eventRemovedViewers)
	case expectedSig("ResetViewers(address,string,string)"):
		eventResetViewers := ResetViewers{}
		errUnpack = eventABI.UnpackIntoInterface(&eventResetViewers, "ResetViewers", vLog.Data)
		errGenerate = generateResetViewersMsg(eventResetViewers)
	case expectedSig("ChangedOwner(address,string,string,string)"):
		eventChangedOwner := ChangedOwner{}
		errUnpack = eventABI.UnpackIntoInterface(&eventChangedOwner, "ChangedOwner", vLog.Data)
		errGenerate = generateChangedOwnerMsg(eventChangedOwner)
	case expectedSig("AddedEditors(address,string,string,string,string)"):
		eventAddedEditors := AddedEditors{}
		errUnpack = eventABI.UnpackIntoInterface(&eventAddedEditors, "AddedEditors", vLog.Data)
		errGenerate = generateAddedEditorsMsg(eventAddedEditors)
	case expectedSig("RemovedEditors(address,string,string,string)"):
		eventRemovedEditors := RemovedEditors{}
		errUnpack = eventABI.UnpackIntoInterface(&eventRemovedEditors, "RemovedEditors", vLog.Data)
		errGenerate = generateRemovedEditorsMsg(eventRemovedEditors)
	case expectedSig("ResetEditors(address,string,string)"):
		eventResetEditors := ResetEditors{}
		errUnpack = eventABI.UnpackIntoInterface(&eventResetEditors, "ResetEditors", vLog.Data)
		errGenerate = generateResetEditorsMsg(eventResetEditors)
	case expectedSig("CreatedNotification(address,string,string,string)"):
		eventCreatedNotification := CreatedNotification{}
		errUnpack = eventABI.UnpackIntoInterface(&eventCreatedNotification, "CreatedNotification", vLog.Data)
		errGenerate = generateCreatedNotificationMsg(eventCreatedNotification)
	case expectedSig("DeletedNotification(address,string,uint64)"):
		eventDeletedNotification := DeletedNotification{}
		errUnpack = eventABI.UnpackIntoInterface(&eventDeletedNotification, "DeletedNotification", vLog.Data)
		errGenerate = generateDeletedNotificationMsg(eventDeletedNotification)
	case expectedSig("BlockedSenders(address,string[])"):
		eventBlockedSenders := BlockedSenders{}
		errUnpack = eventABI.UnpackIntoInterface(&eventBlockedSenders, "BlockedSenders", vLog.Data)
		errGenerate = generateBlockedSendersMsg(eventBlockedSenders)
	default:
		log.Fatal("Failed to unpack log data into any event type")
	}

	if errUnpack != nil || errGenerate != nil {
		log.Fatalf("Failed to unpack event %v: %v %v", eventSig, errUnpack, errGenerate)
	}

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
		log.Printf("Response is empty")
		return
	}

	log.Println(res.RawLog)
	log.Println(res.TxHash)

	// callback on EVM chain
	account, err := wEth.Derive(hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0"), false)
	if err != nil {
		log.Printf("Failed to derive account: %v", err)
		return
	}

	// generate privkey and command
	privKey, err := wEth.PrivateKeyHex(account)
	if err != nil {
		log.Printf("Failed to generate privkey: %v", err)
		return
	}

	// 0 is a placeholder message id
	cmdArgs := []string{"send", "--rpc-url", RPC, "--private-key", privKey, vLog.Address.Hex(), "finishMessage(string)", "0"}
	cmd := exec.Command("cast", cmdArgs...)
	log.Printf("Executing: %v", cmd.String())

	// Capture debugging outpot
	var stdoutBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	// err = cmd.Run()

	// Extract the transaction hash if present
	output := stdoutBuf.String()
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "transactionHash") {
			log.Printf("Tx: %v", strings.TrimSpace(strings.Fields(line)[1]))
			break
		}
	}
	log.Print("Mock execution complete")
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

func expectedSig(function string) string {
	return crypto.Keccak256Hash([]byte(function)).Hex()
}
