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
	log.Printf("Now listening to %s...", network.Name)

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
		rpcClient, err := ethclient.Dial(network.RPC)
		if err != nil {
			log.Printf("Failed to connect to the Ethereum RPC client, retrying in 5 seconds: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		wsClient, err := ethclient.Dial(network.WS)
		if err != nil {
			log.Printf("Failed to connect to the Ethereum WS client, retrying in 5 seconds: %v", err)

			if wsClient != nil {
				wsClient.Close()
			}
			time.Sleep(5 * time.Second)
			continue
		}

		sub, logs, err = subscribeLogs(wsClient, query)
		if err != nil {
			log.Printf("Failed to subscribe, retrying in 5 seconds: %v", err)
			if wsClient != nil {
				wsClient.Close()
			}
			time.Sleep(5 * time.Second)
			continue
		}

		log.Print("Ready to listen!")

		// Listening loop
		func() {
			defer func() {
				// Unsubscribe and close client on exit
				if sub != nil {
					sub.Unsubscribe()
				}
				wsClient.Close()
				wsClient.Close()
			}()

			for {
				select {
				case <-sigs:
					log.Print("Exiting...")
					stopped = true
					return
				case err := <-sub.Err():
					log.Printf("Subscription error, reconnecting: %v", err)
					return // Break out of the loop to retry
				case vLog := <-logs:
					log.Printf("Log received: %s", vLog.Address.Hex())

					go func(vLog types.Log) {
						err := waitForReceipt(rpcClient, vLog.TxHash, network.ChainID, network.Finality, func(receipt *types.Receipt) {
							for _, l := range receipt.Logs {
								if l.Address.Hex() == contractAddress.Hex() {
									handleLog(l, a.w, a.q, network.ChainID, jackalContract)
								}
							}
						})
						if err != nil {
							log.Printf("Error getting receipt for tx %s: %v", vLog.TxHash.Hex(), err)
						}
					}(vLog)
				}
			}
		}()
	}

	wg.Done()
}
