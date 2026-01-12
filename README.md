# Ethereum Smart Contract Bytecode Analyzer (GasLens)

A comprehensive Go-based tool that analyzes Ethereum smart contract bytecode to identify gas usage patterns, storage hotspots, loops, and function-level gas consumption. This tool helps developers optimize smart contracts for gas efficiency.

## Features

### ✅ Completed Features

1. **Opcode Analysis**
   - Count occurrences of each opcode
   - Sum total gas per opcode
   - Print each opcode with its approximate gas cost

2. **PUSH Instructions**
   - Decode PUSH1–PUSH32 values
   - Print pushed data in hex for readability

3. **Storage Analysis**
   - Track SSTORE (writes) per slot
   - Track SLOAD (reads) per slot
   - Detect consecutive SSTOREs and suggest packing variables
   - Summarize storage hotspots

4. **Stack Simulation**
   - Simulate basic stack operations (POP, DUP, SWAP, PUSH)
   - Track storage slot reads/writes in coordination with the stack

5. **Loop Detection**
   - Track backward jumps using a LoopTracker
   - Record start and end PC of loops
   - Suggest optimizing or limiting iterations

6. **Function-Level Analysis**
   - Track functions using PUSH4 selectors
   - Track gas consumed per function
   - Show top N most expensive functions

7. **Reporting & Visualization**
   - Print opcode frequency summary
   - Print top N expensive opcodes
   - Print storage read/write hotspots
   - Print loop warnings
   - Print function gas summary
   - ASCII bar charts for gas consumption visualization

8. **Export Features**
   - Export analysis reports to JSON format
   - Export analysis reports to CSV format
   - Generate automatic optimization suggestions

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd gaslens
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the project:
```bash
go build -o gaslens main.go
```

## Usage

### Analyze Local Bytecode File

```bash
./gaslens <bytecode_file>
```

Example:
```bash
./gaslens bytecode.txt
```

### Analyze Deployed Contract

Set your Etherscan API key in a `.env` file:
```
ETHERSCAN_API_KEY=your_api_key_here
```

Then analyze a deployed contract:
```bash
./gaslens -address <contract_address>
```

Example:
```bash
./gaslens -address 0x1234567890123456789012345678901234567890
```

## Output

The analyzer provides comprehensive output including:

1. **Detailed Opcode Trace**: Each instruction with its gas cost
2. **Opcode Frequency Summary**: Count and total gas per opcode
3. **Top Expensive Opcodes**: Ranked list of gas-consuming operations
4. **Visual Gas Charts**: ASCII bar charts showing gas distribution
5. **Storage Analysis**: Hotspots for reads and writes
6. **Loop Detection**: Backward jumps and optimization warnings
7. **Function Analysis**: Gas consumption per function selector
8. **Optimization Suggestions**: Automated recommendations

### Export Files

The analyzer automatically generates:
- `analysis_report.json`: Complete analysis in JSON format
- `analysis_report.csv`: Tabular data for spreadsheet analysis

## Example Output

```
=== Top 5 Expensive Opcodes ===
1. SSTORE     : 20000 gas
2. SLOAD      : 200 gas
3. PUSH1      : 51 gas
4. JUMPI      : 50 gas
5. JUMP       : 24 gas

=== Top 5 Most Expensive Functions ===
1. Function 0x2e1a7d4d at PC 32 used approx 20401 gas
2. Function 0x8da5cb5b at PC 43 used approx 20379 gas

=== Optimization Suggestions ===
1. Cache storage slot 0 in memory (read 3 times)
2. High SSTORE usage detected - consider struct packing
```

## Project Structure

```
gaslens/
├── main.go                 # Entry point
├── analyzer/
│   ├── analyzer.go         # Main analysis engine
│   ├── gas_table.go        # EVM opcode gas costs
│   ├── storage.go          # Storage tracking
│   ├── stack.go            # Stack simulation
│   ├── loop.go             # Loop detection
│   ├── function_tracker.go # Function analysis
│   └── reporter.go         # Export and reporting
├── utils/
│   ├── file.go             # File operations
│   └── etherscan.go        # Etherscan API integration
├── bytecode.txt            # Sample bytecode
└── README.md
```

## Gas Cost Table

The analyzer uses approximate gas costs based on the Ethereum Yellow Paper:

- **SSTORE**: 20,000 gas (storage write)
- **SLOAD**: 100 gas (storage read)
- **CALL**: 700 gas (external call)
- **JUMPI**: 10 gas (conditional jump)
- **PUSH1-32**: 3 gas (push operations)
- And many more...

## Optimization Suggestions

The tool automatically suggests optimizations such as:

1. **Storage Caching**: When slots are read multiple times
2. **Struct Packing**: When consecutive SSTORE operations are detected
3. **Loop Optimization**: When backward jumps are frequent
4. **Gas Limit Warnings**: For potentially expensive operations

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Built using the go-ethereum library for EVM opcodes
- Inspired by gas optimization needs in smart contract development
- Uses Etherscan API for deployed contract analysis
