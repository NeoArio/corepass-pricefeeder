package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NeoArio/corepass-pricefeeder/contracts"
	"github.com/core-coin/go-core/accounts/abi/bind"
	"github.com/core-coin/go-core/common"
	"github.com/core-coin/go-core/xcbclient"
	//"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"net/http"
)

func GetSimpleStorageCoreCoin(w http.ResponseWriter, req *http.Request) {
	//client, err := xcbclient.Dial("https://stg.pingextest.eu/xcb")
	client, err := xcbclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	//Invalid network id prefix in address
	common.DefaultNetworkID = 3

	address, err := common.HexToAddress("ab81c802519469c525aca73c603699678ff3771ff6a6")
	if err != nil {
		log.Fatal(err)
	}
	instance, err := contracts.NewSimple(address, client)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	opts := &bind.CallOpts{
		From: address,
		Context: ctx,
	}

	storage, err := instance.Get(opts)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(storage) // "1.0"

	json.NewEncoder(w).Encode(storage)
}