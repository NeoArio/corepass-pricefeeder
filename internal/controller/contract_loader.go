package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NeoArio/corepass-pricefeeder/internal/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/NeoArio/corepass-pricefeeder/contracts" // for demo
)


func GetSimpleStorage(w http.ResponseWriter, req *http.Request) {
    client, err := ethclient.Dial("https://kovan.infura.io/v3/8216411c128a4b5480fd2ee0e2771d7b")
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress(core.SimpleStorageAddress)
    instance, err := contracts.NewContracts(address, client)
    if err != nil {
        log.Fatal(err)
    }

	ctx := context.Background()

	opts := &bind.CallOpts{
		Context: ctx,
	}

	version, err := instance.Get(opts)
    if err != nil {
        log.Println(err)
    }

    fmt.Println(version) // "1.0"

	json.NewEncoder(w).Encode(version)
}