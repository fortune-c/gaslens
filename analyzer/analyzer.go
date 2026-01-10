package analyzer

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/core/vm"
	"sort"
)

type opGasPair struct {
	Op  vm.OpCode
	Gas uint64
}

// AnalyzeBytecode prints opcode gas and charts
func AnalyzeBytecode(code []byte) {
	pc := 0
	var totalGas uint64

	loopTracker := NewLoopTracker()
	functionTracker := NewFunctionTracker()
	opcodeCount := map[vm.OpCode]int{}
	opcodeGas := map[vm.OpCode]uint64{}
	engine := &StackEngine{}
	storage := NewStorageTracker()

	var consecutiveSSTORE int
	var maxConsecutiveSSTORE int

	for pc < len(code) {
		currentPC := pc
		op := vm.OpCode(code[pc])
		if op == vm.PUSH4 && pc+5 <= len(code) {
			selectorBytes := code[pc+1 : pc+5]
			selector := fmt.Sprintf("0x%x", selectorBytes)
			functionTracker.AddFunction(selector, pc)
		}
		switch {
		case op >= vm.PUSH1 && op <= vm.PUSH32:
			size := int(op - vm.PUSH1 + 1)
			if pc+1+size <= len(code) {
				valueBytes := code[pc+1 : pc+1+size]
				var v uint64
				for _, b := range valueBytes {
					v = (v << 8) | uint64(b)
				}
				engine.Push(v)
			}

		case op == vm.POP:
			engine.Pop()

		case op >= vm.DUP1 && op <= vm.DUP16:
			n := int(op - vm.DUP1 + 1)
			engine.Dup(n)

		case op >= vm.SWAP1 && op <= vm.SWAP16:
			n := int(op - vm.SWAP1 + 1)
			engine.Swap(n)

		case op == vm.SLOAD:
			slot := engine.Pop()
			storage.SLoadCount[slot]++
			engine.Push(0)

		case op == vm.SSTORE:
			slot := engine.Pop()
			_ = engine.Pop()
			storage.SStoreCount[slot]++
		}
		opcodeCount[op]++
		gas := GetGasCost(op)
		totalGas += gas
		opcodeGas[op] += gas
		functionTracker.AddGas(currentPC, gas)
		if op == vm.SSTORE {
			consecutiveSSTORE++
			if consecutiveSSTORE > maxConsecutiveSSTORE {
				maxConsecutiveSSTORE = consecutiveSSTORE
			}
		} else {
			consecutiveSSTORE = 0
		}

		fmt.Printf("%04d: %-10s Gas: %d\n", pc, op.String(), gas)
		pc++
		if op == vm.JUMP || op == vm.JUMPI {
			if currentPC > 0 {
				prevOp := vm.OpCode(code[currentPC-1])

				if prevOp >= vm.PUSH1 && prevOp <= vm.PUSH32 {
					pushSize := int(prevOp - vm.PUSH1 + 1)

					destStart := currentPC - pushSize
					if destStart >= 0 {
						destBytes := code[destStart:currentPC]
						dest := 0
						for _, b := range destBytes {
							dest = (dest << 8) | int(b)
						}

						if dest < currentPC {
							loopTracker.RecordLoop(dest, currentPC)
						}
					}
				}
			}
		}
		if op >= vm.PUSH1 && op <= vm.PUSH32 {
			pushBytes := int(op - vm.PUSH1 + 1)
			if pc+pushBytes > len(code) {
				break
			}
			value := code[pc : pc+pushBytes]
			fmt.Printf("      PUSH Data: 0x%s\n", hex.EncodeToString(value))
			pc += pushBytes
		}
	}

	if maxConsecutiveSSTORE > 1 {
		fmt.Printf("\nDetected %d consecutive SSTORE instructions. Consider packing variables.\n", maxConsecutiveSSTORE)
	}

	// Summary
	fmt.Println("\n=== Opcode Frequency Summary ===")
	for op, count := range opcodeCount {
		fmt.Printf("%-10s : %d times, approx gas: %d\n", op.String(), count, count*int(GetGasCost(op)))
	}

	fmt.Println("\n=== Top 5 Expensive Opcodes ===")
	for i, pair := range topExpensiveOpcodes(opcodeGas, 5) {
		fmt.Printf("%d. %-10s : %d gas\n", i+1, pair.Op.String(), pair.Gas)
	}

	fmt.Println("\nTotal Approximate Gas Cost:", totalGas)

	// Charts
	opcodeChart := map[string]uint64{}
	for op, gas := range opcodeGas {
		opcodeChart[op.String()] = gas
	}
	printBarChart("Top Gas-Consuming Opcodes", opcodeChart, 50)

	fmt.Println("\n=== Storage Write Hotspots ===")
	for slot, count := range storage.SStoreCount {
		if count > 1 {
			fmt.Printf("Slot %d written %d times – consider packing or caching\n", slot, count)
		}
	}

	fmt.Println("\n=== Repeated Storage Reads ===")
	for slot, count := range storage.SLoadCount {
		if count > 2 {
			fmt.Printf("Slot %d read %d times – cache in memory variable\n", slot, count)
		}
	}

	fmt.Println("\n=== Loop Detection ===")
	if len(loopTracker.Loops) == 0 {
		fmt.Println("No backward jumps detected (no loops found)")
	} else {
		for _, loop := range loopTracker.Loops {
			fmt.Printf(
				"Loop detected: PC %d -> %d (executed %d times) – consider optimizing or limiting iterations\n",
				loop.StartPC,
				loop.EndPC,
				loop.Count,
			)
		}
	}

	fmt.Println("\n=== Function Gas Usage ===")
	if len(functionTracker.Functions) == 0 {
		fmt.Println("No function selectors detected")
	} else {
		for _, fn := range functionTracker.Functions {
			fmt.Printf("Function %s at PC %d used approx %d gas\n",
				fn.Selector,
				fn.EntryPC,
				fn.Gas,
			)
		}
	}

}

// Helper functions

func topExpensiveOpcodes(opcodeGas map[vm.OpCode]uint64, topN int) []opGasPair {
	pairs := []opGasPair{}
	for op, gas := range opcodeGas {
		pairs = append(pairs, opGasPair{op, gas})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Gas > pairs[j].Gas })
	if len(pairs) > topN {
		return pairs[:topN]
	}
	return pairs
}

func printBarChart(label string, data map[string]uint64, maxWidth int) {
	var maxVal uint64
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}
	if maxVal == 0 {
		fmt.Println("No data for", label)
		return
	}
	fmt.Println("\n=== " + label + " ===")
	for k, v := range data {
		barLen := int(v * uint64(maxWidth) / maxVal)
		bar := ""
		for i := 0; i < barLen; i++ {
			bar += "█"
		}
		fmt.Printf("%-10s |%-*s %d\n", k, maxWidth, bar, v)
	}
}
