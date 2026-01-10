package utils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// FetchBytecode fetches deployed bytecode from Etherscan for a contract address
func FetchBytecode(address, apiKey string) []byte {
	url := fmt.Sprintf("https://api.etherscan.io/v2/api?chainid=1&module=proxy&action=eth_getCode&address=%s&tag=latest&apikey=%s", address, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch bytecode: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	codeHex, ok := result["result"].(string)
	if !ok || codeHex == "0x" {
		log.Fatalf("Invalid contract address or no bytecode found")
	}

	// Remove "0x" prefix if present
	codeHex = strings.TrimPrefix(codeHex, "0x")

	// Ensure lowercase and no whitespace
	codeHex = strings.ToLower(strings.TrimSpace(codeHex))

	// Decode hex string
	code, err := hex.DecodeString(codeHex)
	if err != nil {
		log.Fatalf("Failed to decode hex string: %v", err)
	}

	return code
}
