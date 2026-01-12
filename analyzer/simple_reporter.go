package analyzer

import (
	"fmt"
	"sort"
)

// SimpleReport provides user-friendly analysis
func PrintSimpleReport(report *AnalysisReport) {
	fmt.Println("\n🔍 SMART CONTRACT GAS ANALYSIS")
	fmt.Println("================================")
	
	// Overall gas estimate
	fmt.Printf("💰 Estimated Total Gas Cost: %d gas\n", report.TotalGas)
	fmt.Printf("💵 Approximate Cost (20 gwei): $%.4f USD\n", estimateUSDCost(report.TotalGas))
	
	// Simple gas rating
	rating := getGasRating(report.TotalGas)
	fmt.Printf("⭐ Gas Efficiency Rating: %s\n\n", rating)
	
	// Top 3 most expensive operations (simplified)
	fmt.Println("🔥 TOP GAS CONSUMERS:")
	printTopOperations(report.TopExpensiveOps, 3)
	
	// Storage efficiency
	fmt.Println("\n💾 STORAGE USAGE:")
	printStorageEfficiency(report.StorageReads, report.StorageWrites)
	
	// Simple recommendations
	fmt.Println("\n💡 OPTIMIZATION TIPS:")
	printSimpleOptimizations(report.Optimizations)
	
	// Function costs (if any)
	if len(report.Functions) > 0 {
		fmt.Println("\n🎯 FUNCTION COSTS:")
		printFunctionCosts(report.Functions, 3)
	}
}

func estimateUSDCost(gas uint64) float64 {
	// Rough estimate: 20 gwei * gas * $3000 ETH price
	gweiCost := float64(gas) * 20 / 1e9 // Convert to ETH
	return gweiCost * 3000 // Rough ETH price
}

func getGasRating(gas uint64) string {
	switch {
	case gas < 50000:
		return "🟢 EXCELLENT (Very efficient)"
	case gas < 100000:
		return "🟡 GOOD (Moderately efficient)"
	case gas < 200000:
		return "🟠 FAIR (Could be optimized)"
	default:
		return "🔴 POOR (Needs optimization)"
	}
}

func printTopOperations(ops []OpGasPair, limit int) {
	if len(ops) == 0 {
		fmt.Println("   No expensive operations found")
		return
	}
	
	for i, op := range ops {
		if i >= limit {
			break
		}
		
		description := getOperationDescription(op.Opcode)
		fmt.Printf("   %d. %s - %d gas\n", i+1, description, op.Gas)
	}
}

func getOperationDescription(opcode string) string {
	descriptions := map[string]string{
		"SSTORE": "💾 Storage Write (expensive!)",
		"SLOAD":  "📖 Storage Read",
		"CALL":   "📞 External Call",
		"CREATE": "🏗️  Contract Creation",
		"JUMPI":  "🔀 Conditional Logic",
		"JUMP":   "➡️  Code Navigation",
		"PUSH1":  "📥 Data Loading",
		"PUSH2":  "📥 Data Loading",
		"PUSH4":  "📥 Function Selector",
	}
	
	if desc, exists := descriptions[opcode]; exists {
		return desc
	}
	return fmt.Sprintf("⚙️  %s Operation", opcode)
}

func printStorageEfficiency(reads map[uint64]int, writes map[uint64]int) {
	totalReads := 0
	totalWrites := 0
	
	for _, count := range reads {
		totalReads += count
	}
	for _, count := range writes {
		totalWrites += count
	}
	
	if totalReads == 0 && totalWrites == 0 {
		fmt.Println("   📊 No storage operations detected")
		return
	}
	
	fmt.Printf("   📖 Storage Reads: %d\n", totalReads)
	fmt.Printf("   💾 Storage Writes: %d\n", totalWrites)
	
	if totalWrites > 5 {
		fmt.Println("   ⚠️  High storage writes - consider batching")
	}
	if totalReads > 10 {
		fmt.Println("   ⚠️  Many storage reads - consider caching")
	}
}

func printSimpleOptimizations(optimizations []string) {
	if len(optimizations) == 0 {
		fmt.Println("   ✅ No obvious optimizations needed!")
		return
	}
	
	simplified := make(map[string]string)
	simplified["Cache storage"] = "💡 Store frequently used data in memory instead of storage"
	simplified["High SSTORE"] = "💡 Group related variables together to save storage costs"
	simplified["Loop"] = "💡 Add gas limits to loops to prevent failures"
	
	count := 1
	for _, opt := range optimizations {
		for key, simple := range simplified {
			if contains(opt, key) {
				fmt.Printf("   %d. %s\n", count, simple)
				count++
				break
			}
		}
		if count > 3 { // Limit to top 3 suggestions
			break
		}
	}
}

func printFunctionCosts(functions []FunctionInfo, limit int) {
	sorted := make([]FunctionInfo, len(functions))
	copy(sorted, functions)
	
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Gas > sorted[j].Gas
	})
	
	for i, fn := range sorted {
		if i >= limit {
			break
		}
		
		costLevel := "💚 Cheap"
		if fn.Gas > 50000 {
			costLevel = "💛 Moderate"
		}
		if fn.Gas > 100000 {
			costLevel = "❤️ Expensive"
		}
		
		fmt.Printf("   %d. Function %s - %d gas (%s)\n", i+1, fn.Selector, fn.Gas, costLevel)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
