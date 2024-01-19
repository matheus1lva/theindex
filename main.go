package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"indexer/constants"
	"indexer/filters"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/X_Wf4I5eScaHgCbO0uLOcajD6REkEZty")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber, err := client.BlockNumber(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(blockNumber)),
		ToBlock:   big.NewInt(int64(blockNumber)),
		Addresses: []common.Address{
			common.HexToAddress("0x00000000000000ADc04C56Bf30aC9d3c0aAF14dC"),
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)

	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

	for _, vLog := range logs {
		// Check for specific properties in each log
		fmt.Println("Log Address:", vLog.Address.Hex()) // Example: Print the log's address

		fmt.Println("Equal ", vLog.Address.Hex() == constants.SeaportEth, constants.SeaportEth)
		if vLog.Address.Hex() == constants.SeaportEth {
			jsonLogs, err := json.MarshalIndent(vLog, "", "    ")
			if err != nil {
				log.Fatalf("Failed to marshal logs: %v", err)
			}

			fmt.Println(string(jsonLogs))
			go filters.Opensea(vLog)
		}

	}
}
