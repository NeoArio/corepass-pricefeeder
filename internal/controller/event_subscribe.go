package controller

import (
	"context"
	"fmt"
	"github.com/NeoArio/corepass-pricefeeder/internal/core"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func EventSubscribe() {

	client, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/8216411c128a4b5480fd2ee0e2771d7b")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(core.SimpleStorageV2Address)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}