package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NeoArio/corepass-pricefeeder/contracts"
	"github.com/NeoArio/corepass-pricefeeder/internal/core"
	"github.com/core-coin/go-core/accounts/abi/bind"
	"github.com/core-coin/go-core/common"
	"github.com/core-coin/go-core/xcbclient"
	"github.com/spf13/viper"

	//"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetSimpleStorageCoreCoin(w http.ResponseWriter, req *http.Request) {
	//client, err := xcbclient.Dial("https://stg.pingextest.eu/xcb")
	client, err := xcbclient.Dial(viper.GetString("xcb_node_url"))
	if err != nil {
		log.Error(err)
		return
	}
	//Invalid network id prefix in address
	common.DefaultNetworkID = 3

	address, err := common.HexToAddress(core.PriceFeedAddress)
	if err != nil {
		log.Error(err)
		return
	}
	instance, err := contracts.NewPriceFeeder(address, client)
	if err != nil {
		log.Error(err)
		return
	}

	ctx := context.Background()

	opts := &bind.CallOpts{
		From: address,
		Context: ctx,
	}

	v1, v2, err := instance.GetLatestPrice(opts)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debugf("%v, %v", v1, v2) // "1.0"

	json.NewEncoder(w).Encode(fmt.Sprintf("%v, %v", v1, v2))
}