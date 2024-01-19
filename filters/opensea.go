package filters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func fetchABI() (abi.ABI, error) {
	// Define your smart contract address and Etherscan API key
	smartContract := "0x00000000000000ADc04C56Bf30aC9d3c0aAF14dC"
	etherscanAPIKey := "BAN2YFN7S4URKNMFCRRYHM8P6X1Y7PA1I5"

	// Construct the API endpoint URL
	abiEndpoint := fmt.Sprintf("https://api.etherscan.io/api?module=contract&action=getabi&address=%s&apikey=%s", smartContract, etherscanAPIKey)

	// Make the HTTP GET request
	resp, err := http.Get(abiEndpoint)
	if err != nil {
		// Handle error
		fmt.Println("Error making the request:", err)
		return abi.ABI{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle error
		fmt.Println("Error reading the response body:", err)
		return abi.ABI{}, err
	}

	// Define a variable to hold the parsed JSON
	var cabi abi.ABI

	// Unmarshal the JSON into the variable
	err = json.Unmarshal(body, &cabi)

	return cabi, nil

}

func Opensea(log types.Log) {
	// receiptEvent := log.Topics[0].Hex()

	abi, err := fetchABI()

	if err != nil {
		fmt.Println(err)
	}

	// abiEvent := abi.Events[receiptEvent]
	fmt.Println("event ", abi)

	// var decodedData []interface{}
	// err = abi.UnpackIntoInterface(&decodedData, "OrderFulfilled", log.Data)
	// if err != nil {
	// 	fmt.Println("Failed to unpack data: ", err)
	// }

}
