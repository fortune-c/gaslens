package analyzer

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/ethereum/go-ethereum/core/vm"
)

type AnalysisReport struct {
	TotalGas           uint64                    `json:"total_gas"`
	OpcodeFrequency    map[string]int            `json:"opcode_frequency"`
	OpcodeGas          map[string]uint64         `json:"opcode_gas"`
	StorageReads       map[uint64]int            `json:"storage_reads"`
	StorageWrites      map[uint64]int            `json:"storage_writes"`
	Loops              []Loop                    `json:"loops"`
	Functions          []FunctionInfo            `json:"functions"`
	TopExpensiveOps    []OpGasPair               `json:"top_expensive_opcodes"`
	Optimizations      []string                  `json:"optimization_suggestions"`
}

type OpGasPair struct {
	Opcode string `json:"opcode"`
	Gas    uint64 `json:"gas"`
}

func ExportToJSON(report *AnalysisReport, filename string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func ExportToCSV(report *AnalysisReport, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	writer.Write([]string{"Category", "Item", "Value", "Details"})

	// Write opcode data
	for opcode, count := range report.OpcodeFrequency {
		gas := report.OpcodeGas[opcode]
		writer.Write([]string{"Opcode", opcode, strconv.Itoa(count), fmt.Sprintf("Gas: %d", gas)})
	}

	// Write storage data
	for slot, count := range report.StorageReads {
		writer.Write([]string{"Storage Read", fmt.Sprintf("Slot %d", slot), strconv.Itoa(count), ""})
	}

	for slot, count := range report.StorageWrites {
		writer.Write([]string{"Storage Write", fmt.Sprintf("Slot %d", slot), strconv.Itoa(count), ""})
	}

	// Write function data
	for _, fn := range report.Functions {
		writer.Write([]string{"Function", fn.Selector, strconv.FormatUint(fn.Gas, 10), fmt.Sprintf("Entry PC: %d", fn.EntryPC)})
	}

	return nil
}

func GenerateOptimizationSuggestions(storage *StorageTracker, loops []Loop, opcodeGas map[vm.OpCode]uint64) []string {
	var suggestions []string

	// Storage optimization suggestions
	for slot, count := range storage.SLoadCount {
		if count > 3 {
			suggestions = append(suggestions, fmt.Sprintf("Cache storage slot %d in memory (read %d times)", slot, count))
		}
	}

	// Loop optimization suggestions
	if len(loops) > 0 {
		suggestions = append(suggestions, "Consider gas limits for loops to prevent out-of-gas errors")
		for _, loop := range loops {
			if loop.Count > 5 {
				suggestions = append(suggestions, fmt.Sprintf("Loop at PC %d-%d executed %d times - consider optimization", loop.StartPC, loop.EndPC, loop.Count))
			}
		}
	}

	// Expensive opcode suggestions
	if gas, exists := opcodeGas[vm.SSTORE]; exists && gas > 100000 {
		suggestions = append(suggestions, "High SSTORE usage detected - consider struct packing")
	}

	if gas, exists := opcodeGas[vm.SLOAD]; exists && gas > 50000 {
		suggestions = append(suggestions, "High SLOAD usage detected - cache frequently accessed storage")
	}

	return suggestions
}

func GetTopExpensiveFunctions(functions []FunctionInfo, topN int) []FunctionInfo {
	sorted := make([]FunctionInfo, len(functions))
	copy(sorted, functions)
	
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Gas > sorted[j].Gas
	})
	
	if len(sorted) > topN {
		return sorted[:topN]
	}
	return sorted
}
