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
		fmt.Println("Usage:")
		fmt.Println("  gaslens <bytecode_file>                    # Simple analysis")
		fmt.Println("  gaslens -address <contract_address>        # Analyze deployed contract")
		fmt.Println("  gaslens -detailed <bytecode_file>          # Detailed technical analysis")
		return
	}

	var code []byte
	detailed := false

	// Check for flags
	if os.Args[1] == "-address" && len(os.Args) >= 3 {
		apiKey := os.Getenv("ETHERSCAN_API_KEY")
		if apiKey == "" {
			log.Fatal("ETHERSCAN_API_KEY not set. Please set it in your environment or .env file")
		}

		address := os.Args[2]
		code = utils.FetchBytecode(address, apiKey)
	} else if os.Args[1] == "-detailed" && len(os.Args) >= 3 {
		detailed = true
		code = utils.ReadHexFile(os.Args[2])
	} else {
		code = utils.ReadHexFile(os.Args[1])
	}

	analyzer.AnalyzeBytecode(code, detailed)
}
