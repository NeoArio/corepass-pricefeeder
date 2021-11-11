package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NeoArio/corepass-pricefeeder/internal/model"
	"github.com/core-coin/go-core/accounts/abi/bind"
	"github.com/core-coin/go-core/crypto"
	eddsa "github.com/core-coin/go-goldilocks"
	"github.com/sethgrid/pester"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"github.com/core-coin/go-core/common"
	"github.com/core-coin/go-core/xcbclient"

	store "github.com/NeoArio/corepass-pricefeeder/contracts"
)


func AddPrice() {
	client, err := xcbclient.Dial(viper.GetString("XCB_NODE_URL"))
	if err != nil {
		log.Errorf("error in Dialing to xcb node: %v", err)
		return
	}

	common.DefaultNetworkID = common.NetworkID(viper.GetInt("network_id"))

	privateKey, err := getPrivateKey()
	if err != nil {
		log.Errorf("error in reading private key: %v", err)
		return
	}

	fromAddress, err := computeFromAddress(privateKey)
	if err != nil{
		log.Errorf("error in computeFromAddress: %v", err)
		return
	}

	signer, err := NewSigner(privateKey, client, fromAddress)
	if err != nil {
		log.Errorf("error in creating signer: %v", err)
		return
	}

	instance, err := NewPriceFeederInstance(client)
	if err != nil{
		log.Errorf("error in creating contract instance: %v", err)
		return
	}

	price, err := GetLatestPrice()
	if err != nil{
		log.Error(err)
		return
	}

	bigIntPrice := ConvertFloatToBigInt(price)

	transaction, err := instance.AddPrice(signer, bigIntPrice)
	if err != nil {
		log.Errorf("error in adding price: %v", err)
		return
	}

	log.Infof("%+v\n", *transaction)
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
		return nil, err
	}

	energyPrice, err := client.SuggestEnergyPrice(context.Background())
	if err != nil {
		log.Printf("error in getting suggested energy price: %v", err)
		return nil, err
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
	address, err := common.HexToAddress(viper.GetString("PRICE_FEED_ADDRESS"))
	if err != nil {
		return nil, err
	}

	instance, err = store.NewPriceFeeder(address, client)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func GetLatestPrice() (price float64, err error){
	APIAddress := viper.GetString("CTN_PRICE_API_ADDRESS")

	startDate := time.Now().AddDate(-1,0,0).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")

	urlParams := url.Values{"startDate": {startDate}, "endDate": {endDate}, "count": {"1"}}

	client := pester.New()
	client.Concurrency = 1
	client.MaxRetries = 0
	client.Timeout = 10 * time.Second

	req, _ := http.NewRequest("GET", APIAddress + "?" + urlParams.Encode(), nil)
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var CTNTrades model.CTNTrades
	if err := json.NewDecoder(resp.Body).Decode(&CTNTrades); err != nil {
		return 0, err
	}

	if len(CTNTrades.Data)>0{
		price = CTNTrades.Data[0].Close
		return price, nil
	} else{
		return 0, errors.New("response is empty")
	}
}

func ConvertFloatToBigInt(number float64) *big.Int{
	bigInt := big.NewInt(int64(number * 1000000))
	bigInt.Mul(bigInt, big.NewInt(1000000000000))
	return bigInt
}