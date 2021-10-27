package controller

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/NeoArio/corepass-pricefeeder/internal/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "github.com/NeoArio/corepass-pricefeeder/contracts" // for demo
)


func SetSimpleStorageCoreCoin(w http.ResponseWriter, req *http.Request) {
	client, err := ethclient.Dial("https://stg.pingextest.eu/xcb")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("cdac8f92ec2b559ba5821c1e52b6781f73e139b02aab2645a8ebdc0af506497c")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}


	chainID, err := client.ChainID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice



	address := common.HexToAddress(core.SimpleStorageAddress)
	instance, err := store.NewContracts(address, client)
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