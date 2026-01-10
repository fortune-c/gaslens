package main

import (
	"fmt"
	"log"
	"os"

	"gaslens/analyzer"
	"gaslens/utils"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: gaslens <bytecode_file> OR gaslens -address <contract_address>")
		return
	}

	var code []byte

	if os.Args[1] == "-address" && len(os.Args) >= 3 {
		apiKey := os.Getenv("ETHERSCAN_API_KEY")
		if apiKey == "" {
			log.Fatal("ETHERSCAN_API_KEY not set. Please set it in your environment or .env file")
		}

		address := os.Args[2]
		code = utils.FetchBytecode(address, apiKey) // Fetch deployed contract bytecode
	} else {
		code = utils.ReadHexFile(os.Args[1]) // Read local bytecode file
	}

	analyzer.AnalyzeBytecode(code)
}
