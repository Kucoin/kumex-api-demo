package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kumex-api-test/sdk"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	kumex *sdk.ApiService
	PATH  = "./config.json.example"
)

func main() {
	//1. get api authentication info from user input or configuration file.
	var apiURI, apiKey, apiSecret, apiPassphrase string
	var apiSkipVerifyTls bool

	// read config from json file
	file, err := ioutil.ReadFile(PATH)
	if err != nil {
		log.Fatal("error when read configuration from json file:", err)
		return
	}
	var config ApiConfig
	// unmarshal config data
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("error when unmarshal config from json file:", err)
		return
	}
	// if we have configuration from json file
	if len(config.ApiBaseURI) > 0 {
		apiURI = config.ApiBaseURI
	} else {
		log.Println("please input api base URI,such as:https://api.kumex.com")
		fmt.Scanf("%s", &apiURI)
	}

	if len(config.ApiKey) > 0 {
		apiURI = config.ApiKey
	} else {
		log.Println("please input your api key...")
		fmt.Scanf("%s", &apiKey)
	}

	if len(config.ApiSecret) > 0 {
		apiURI = config.ApiSecret
	} else {
		log.Println("please input your api secret...")
		fmt.Scanf("%s", &apiSecret)
	}

	if len(config.ApiPassphrase) > 0 {
		apiURI = config.ApiPassphrase
	} else {
		log.Println("please input your api passphrase...")
		fmt.Scanf("%s", &apiPassphrase)
	}

	apiSkipVerifyTls = config.ApiSkipVerifyTls

	initApiService(apiURI, apiKey, apiSecret, apiPassphrase, apiSkipVerifyTls)

	// 2. api function map to user input.
	apiMap := initApiMap()
	// 3. get api info from user input.
	var apiName string
	log.Println("please input the api name which you want to test...")
	fmt.Scanf("%s", &apiName)
	log.Println("please input the api parameters,use ',' to spilt parameters.")
	log.Println("if it is no parameter api,just enter to skip...")

	reader := bufio.NewReader(os.Stdin)
	inputs, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("input parameter error:", err)
	}

	var params []string
	if len(strings.TrimSpace(inputs)) > 0 {
		params = strings.Split(strings.TrimSpace(inputs), ",")
	}
	// 4. call api function which user choose to test.
	call(apiMap, apiName, params)
}

type ApiConfig struct {
	ApiBaseURI       string `json:ApiBaseURI`
	ApiKey           string `json:ApiKey`
	ApiSecret        string `json:ApiSecret`
	ApiPassphrase    string `json:ApiPassphrase`
	ApiSkipVerifyTls bool   `json:ApiSkipVerifyTls`
}

// initApiService init an api service
func initApiService(apiURI, apiKey, apiSecret, apiPassphrase string, apiSkipVerifyTls bool) {
	kumex = sdk.NewApiService(
		sdk.ApiBaseURIOption(apiURI),
		sdk.ApiKeyOption(apiKey),
		sdk.ApiSecretOption(apiSecret),
		sdk.ApiPassPhraseOption(apiPassphrase),
		sdk.ApiSkipVerifyTlsOption(apiSkipVerifyTls),
	)
	sdk.DebugMode = true
}

// call according to string input to call function which named the input.
func call(m map[string]interface{}, name string, params []string) {
	// if string input is a function
	f := reflect.ValueOf(m[name])
	if f.Kind() != reflect.Func {
		log.Fatal("input error,your input is not a name of api.")
		return
	}

	in := make([]reflect.Value, len(params))
	// if the function has no parameters
	if len(params) > 0 {
		for k, v := range params {
			fmt.Println(v)
			in[k] = reflect.ValueOf(v)
		}
	} else {
		in = nil
	}

	f.Call(in)
}

// initApiMap init api function mapping
func initApiMap() map[string]interface{} {
	return map[string]interface{}{
		"AccountOverview":    kumex.AccountOverview,
		"TransactionHistory": kumex.TransactionHistory,

		"DepositAddresses": kumex.DepositAddresses,
		"Deposits":         kumex.Deposits,
		"WithdrawalQuotas": kumex.WithdrawalQuotas,
		"ApplyWithdrawal":  kumex.ApplyWithdrawal,
		"Withdrawals":      kumex.Withdrawals,
		"CancelWithdrawal": kumex.CancelWithdrawal,

		"TransferOut":    kumex.TransferOut,
		"TransferOutV2":  kumex.TransferOutV2,
		"TransferList":   kumex.TransferList,
		"CancelTransfer": kumex.CancelTransfer,

		"Fills":               kumex.Fills,
		"RecentFills":         kumex.RecentFills,
		"OpenOrderStatistics": kumex.OpenOrderStatistics,

		"CreateOrder":      kumex.CreateOrder,
		"CancelOrder":      kumex.CancelOrder,
		"CancelOrders":     kumex.CancelOrders,
		"StopOrders":       kumex.StopOrders,
		"GetStopOrders":    kumex.GetStopOrders,
		"Orders":           kumex.Orders,
		"Order":            kumex.Order,
		"RecentDoneOrders": kumex.RecentDoneOrders,

		"Ticker":             kumex.Ticker,
		"Level2Snapshot":     kumex.Level2Snapshot,
		"Level2MessageQuery": kumex.Level2MessageQuery,
		"Level3Snapshot":     kumex.Level3Snapshot,
		"Level3MessageQuery": kumex.Level3MessageQuery,
		"TradeHistory":       kumex.TradeHistory,
		"InterestQuery":      kumex.InterestQuery,
		"IndexQuery":         kumex.IndexQuery,
		"MarkPrice":          kumex.MarkPrice,
		"PremiumQuery":       kumex.PremiumQuery,
		"FundingRate":        kumex.FundingRate,
		"ActiveContracts":    kumex.ActiveContracts,
		"Contracts":          kumex.Contracts,

		"WebSocketPublicToken":  kumex.WebSocketPublicToken,
		"WebSocketPrivateToken": kumex.WebSocketPrivateToken,
		"NewWebSocketClient":    kumex.NewWebSocketClient,
	}
}
