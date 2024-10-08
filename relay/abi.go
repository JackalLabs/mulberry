package relay

import (
	"context"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/JackalLabs/mulberry/jackal/uploader"
	evmTypes "github.com/JackalLabs/mulberry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/cosmos-go-wallet/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var ABI = `[
    {
      "type": "function",
      "name": "postFile",
      "inputs": [
        {
          "name": "merkle",
          "type": "string",
          "internalType": "string"
        },
        {
          "name": "filesize",
          "type": "uint64",
          "internalType": "uint64"
        }
      ],
      "outputs": [],
      "stateMutability": "payable"
    },
    {
      "type": "event",
      "name": "PostedFile",
      "inputs": [
        {
          "name": "sender",
          "type": "address",
          "indexed": false,
          "internalType": "address"
        },
        {
          "name": "merkle",
          "type": "string",
          "indexed": false,
          "internalType": "string"
        },
        {
          "name": "size",
          "type": "uint64",
          "indexed": false,
          "internalType": "uint64"
        }
      ],
      "anonymous": false
    }
  ]`

var eventABI abi.ABI

func init() {
	var err error
	eventABI, err = abi.JSON(strings.NewReader(ABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}
}

func handleLog(vLog *types.Log, w *wallet.Wallet, q *uploader.Queue, chainID uint64, jackalContract string) {
	event := struct {
		Sender common.Address
		Merkle string
		Size   uint64
	}{}

	e, err := eventABI.Unpack("PostedFile", vLog.Data)
	if err != nil {
		log.Fatalf("Failed to unpack log data: %v", err)
		return
	}

	fmt.Println(len(e))

	err = eventABI.UnpackIntoInterface(&event, "PostedFile", vLog.Data)
	if err != nil {
		log.Fatalf("Failed to unpack log data: %v", err)
		return
	}

	log.Printf("Event details: %+v", event)

	evmAddress := event.Sender.String()

	merkleRoot, err := hex.DecodeString(event.Merkle)
	if err != nil {
		log.Fatalf("Failed to decode merkle: %v", err)
		return
	}

	abci, err := w.Client.RPCClient.ABCIInfo(context.Background())
	if err != nil {
		log.Fatalf("Failed to query ABCI: %v", err)
		return
	}

	log.Printf("Relaying for %s\n", event.Sender.String())

	var hours int64 = 100 * 365 * 24

	merkleBase64 := base64.StdEncoding.EncodeToString(merkleRoot)

	var maxProofs int64 = 3
	fileSize := int64(event.Size)
	storageMsg := evmTypes.ExecuteMsg{
		PostFile: &evmTypes.ExecuteMsgPostFile{
			Merkle:        merkleBase64,
			FileSize:      fileSize,
			ProofInterval: 3600,
			ProofType:     0,
			MaxProofs:     maxProofs,
			Note:          fmt.Sprintf("{\"memo\":\"Relayed from %s for %s\"}", chainRep(chainID), evmAddress),
			Expires:       abci.Response.LastBlockHeight + ((hours * 60 * 60) / 6),
		},
	}

	factoryMsg := evmTypes.ExecuteFactoryMsg{
		CallBindings: &evmTypes.ExecuteMsgCallBindings{
			EvmAddress: &evmAddress,
			Msg:        &storageMsg,
		},
	}

	cost := q.GetCost(fileSize*maxProofs, hours)
	cost = int64(float64(cost) * 1.2)
	c := sdk.NewInt64Coin("ujkl", cost)

	msg := &wasm.MsgExecuteContract{
		Sender:   w.AccAddress(),
		Contract: jackalContract,
		Msg:      factoryMsg.Encode(),
		Funds:    sdk.NewCoins(c),
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
