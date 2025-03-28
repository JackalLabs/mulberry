package relay

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/JackalLabs/mulberry/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (a *App) ListenToEthereumNetwork(network config.NetworkConfig, wg *sync.WaitGroup) {
	subLogger := log.With().Str("network", network.Name).Logger()

	subLogger.Printf("Connecting to %s", network.Name)

	jackalContract := a.cfg.JackalConfig.Contract

	// Specify the contract address
	contractAddress := common.HexToAddress(network.Contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// Subscribe to the logs
	var sub ethereum.Subscription
	var logs chan types.Log

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	stopped := false
	for !stopped {
		_, err := ethclient.Dial(network.RPC)
		if err != nil {
			subLogger.Printf("Failed to connect to the Ethereum RPC client, retrying in 5 seconds: %v", err)
			subLogger.Printf("rpc client: %s", network.RPC)
			time.Sleep(5 * time.Second)
			continue
		}

		wsClient, err := ethclient.Dial(network.WS)
		if err != nil {
			subLogger.Printf("Failed to connect to the Ethereum WS client, retrying in 5 seconds: %v", err)
			subLogger.Printf("ws client: %s", network.WS)

			if wsClient != nil {
				wsClient.Close()
			}
			time.Sleep(5 * time.Second)
			continue
		}

		sub, logs, err = subscribeLogs(wsClient, query)
		if err != nil {
			subLogger.Printf("Failed to subscribe, retrying in 5 seconds: %v", err)
			subLogger.Printf("ws client: %s", network.WS)
			if wsClient != nil {
				wsClient.Close()
			}
			time.Sleep(5 * time.Second)
			continue
		}
		subLogger.Printf("Ready to listen on %s", network.Name)

		// Listening loop
		func() {
			defer func() {
				// Unsubscribe and close client on exit
				if sub != nil {
					sub.Unsubscribe()
				}
				wsClient.Close()
			}()

			for {
				select {
				case <-sigs:
					log.Print("Exiting...")
					stopped = true
					return
				case err := <-sub.Err():
					subLogger.Printf("Subscription error, reconnecting: %v", err)
					return // Break out of the loop to retry
				case vLog := <-logs:
					subLogger.Printf("Log received: %s", vLog.Address.Hex())

					go func(vLog types.Log) {
						err := waitForReceipt(wsClient, vLog.TxHash, network, func(receipt *types.Receipt) {
							for _, l := range receipt.Logs {
								if l.Address.Hex() == contractAddress.Hex() && len(l.Data) > 0 {
									handleLog(l, a.w, a.wEth, a.q, network.ChainID, network.RPC, jackalContract, a.cfg.MulberrySettings.CastPath)
								}
							}
						})
						if err != nil {
							subLogger.Printf("Error getting receipt for tx %s: %v", vLog.TxHash.Hex(), err)
						}
					}(vLog)
				}
			}
		}()
	}

	wg.Done()
}
