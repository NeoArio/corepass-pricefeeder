package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NeoArio/corepass-pricefeeder/internal/core"
	"github.com/core-coin/go-core/accounts/abi/bind"
	"github.com/core-coin/go-core/crypto"
	eddsa "github.com/core-coin/go-goldilocks"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/core-coin/go-core/common"
	"github.com/core-coin/go-core/xcbclient"

	store "github.com/NeoArio/corepass-pricefeeder/contracts"
)


func SetSimpleStorageCoreCoin(w http.ResponseWriter, req *http.Request) {
	client, err := xcbclient.Dial(viper.GetString("xcb_node_url"))
	if err != nil {
		log.Printf("error in Dialing to xcb node: %v", err)
		return
	}

	common.DefaultNetworkID = common.NetworkID(viper.GetInt("network_id"))

	privateKey, err := getPrivateKey()
	if err != nil {
		log.Printf("error in reading private key: %v", err)
		return
	}

	fromAddress, err := computeFromAddress(privateKey)
	if err != nil{
		log.Printf("error in computeFromAddress: %v", err)
		return
	}

	signer, err := NewSigner(privateKey, client, fromAddress)
	if err != nil {
		log.Printf("error in creating signer: %v", err)
		return
	}

	instance, err := NewPriceFeederInstance(client)
	if err != nil{
		log.Printf("error in creating contract instance: %v", err)
		return
	}

	price := GetLatestPrice(req)

	transaction, err := instance.AddPrice(signer, price)
	if err != nil {
		log.Printf("error in adding price: %v", err)
		return
	}

	fmt.Printf("%+v\n", *transaction) // "1.0"
	json.NewEncoder(w).Encode(transaction)
}

func computeFromAddress(privateKey *eddsa.PrivateKey) (fromAddress common.Address, err error){
	pub := eddsa.Ed448DerivePublicKey(*privateKey)
	fromAddress = crypto.PubkeyToAddress(pub)
	return fromAddress, nil
}

func getPrivateKey() (privateKey *eddsa.PrivateKey, err error){
	privateKeyBytes := []byte(viper.GetString("private_key_bytes"))
	if len(privateKeyBytes) == 0{
		privateKeyBytes, err = ioutil.ReadFile("ab832c70acf1a84bb8b2b5c06f67d21c8f31b056361c.wallet")
		if err != nil {
			return nil, err
		}
	}

	privateKey, err = crypto.ToEDDSA(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func NewSigner(privateKey *eddsa.PrivateKey, client *xcbclient.Client, fromAddress common.Address) (signer *bind.TransactOpts, err error){
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("error in computing nonce: %v", err)
		return
	}

	energyPrice, err := client.SuggestEnergyPrice(context.Background())
	if err != nil {
		log.Printf("error in getting suggested energy price: %v", err)
		return
	}

	signer = bind.NewKeyedTransactor(privateKey)
	if err != nil {
		return nil, err
	}

	signer.Nonce = big.NewInt(int64(nonce))
	signer.Value = big.NewInt(0)     // in wei
	signer.EnergyLimit = uint64(300000) // in units
	signer.EnergyPrice = energyPrice

	return signer, nil
}

func NewPriceFeederInstance(client *xcbclient.Client)(instance *store.PriceFeeder, err error){
	address, err := common.HexToAddress(core.PriceFeedAddress)
	if err != nil {
		return nil, err
	}

	instance, err = store.NewPriceFeeder(address, client)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func GetLatestPrice(req *http.Request) *big.Int{
	n := req.FormValue("number")
	number, _ := strconv.Atoi(n)

	price := big.NewInt(int64(number))

	return price
}