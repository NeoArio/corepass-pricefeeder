package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NeoArio/corepass-pricefeeder/internal/core"
	"github.com/core-coin/go-core/accounts/abi/bind"
	"github.com/core-coin/go-core/crypto"
	eddsa "github.com/core-coin/go-goldilocks"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/core-coin/go-core/common"
	"github.com/core-coin/go-core/xcbclient"

	store "github.com/NeoArio/corepass-pricefeeder/contracts" // for demo
)


func SetSimpleStorageCoreCoin(w http.ResponseWriter, req *http.Request) {

	content, err := ioutil.ReadFile("/home/rasool/Downloads/ab832c70acf1a84bb8b2b5c06f67d21c8f31b056361c.wallet")

	if err != nil {
		log.Fatal(err)
	}

	//private := string(content)

	client, err := xcbclient.Dial("https://stg.pingextest.eu/xcb")
	if err != nil {
		log.Fatal(err)
	}


	common.DefaultNetworkID = 3

	//crypto.HexToECDSA(content)
	//privateKey, err := crypto.HexToECDSA(private)
	privateKey, err := crypto.ToEDDSA(content)
	if err != nil {
		log.Fatal(err)
	}

	pub := eddsa.Ed448DerivePublicKey(*privateKey)



	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := pub.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}

	fromAddress := crypto.PubkeyToAddress(pub)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestEnergyPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}


	//chainID, err := client.ChainID(context.Background())
	auth := bind.NewKeyedTransactor(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.EnergyLimit = uint64(300000) // in units
	auth.EnergyPrice = gasPrice



	address, err := common.HexToAddress(core.SimpleStorageCoreAddress)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := store.NewSimple(address, client)
	if err != nil {
		log.Fatal(err)
	}

	n := req.FormValue("number")
	number, _ := strconv.Atoi(n)

	//ctx := req.Context()

	//opts := &bind.TransactOpts{
	//	Context: ctx,
	//}

	storage := big.NewInt(int64(number))

	version, err := instance.Set(auth, storage)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(version) // "1.0"

	json.NewEncoder(w).Encode(version)
}