package utils

import (
	"encoding/hex"
	"io/ioutil"
	"log"
)

// ReadHexFile reads a file and decodes hex, removing 0x prefix if present.
func ReadHexFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	hexStr := string(data)
	if len(hexStr) >= 2 && hexStr[0:2] == "0x" {
		hexStr = hexStr[2:]
	}

	code, err := hex.DecodeString(hexStr)
	if err != nil {
		log.Fatalf("Failed to decode hex: %v", err)
	}

	return code
}
