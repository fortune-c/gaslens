package analyzer

import "github.com/ethereum/go-ethereum/core/vm"

// GasTable is the approximate gas cost of EVM opcodes
var GasTable = map[vm.OpCode]uint64{
	vm.STOP: 0, vm.ADD: 3, vm.MUL: 5, vm.SUB: 3, vm.DIV: 5, vm.SDIV: 5, vm.MOD: 5, vm.SMOD: 5, vm.ADDMOD: 8, vm.MULMOD: 8,
	vm.EXP: 10, vm.SIGNEXTEND: 5, vm.LT: 3, vm.GT: 3, vm.SLT: 3, vm.SGT: 3, vm.EQ: 3, vm.ISZERO: 3, vm.AND: 3, vm.OR: 3,
	vm.XOR: 3, vm.NOT: 3, vm.BYTE: 3, vm.SHL: 3, vm.SHR: 3, vm.SAR: 3, vm.KECCAK256: 30,
	vm.ADDRESS: 2, vm.BALANCE: 100, vm.ORIGIN: 2, vm.CALLER: 2, vm.CALLVALUE: 2, vm.CALLDATALOAD: 3, vm.CALLDATASIZE: 2,
	vm.CALLDATACOPY: 3, vm.CODESIZE: 2, vm.CODECOPY: 3, vm.GASPRICE: 2, vm.EXTCODESIZE: 100, vm.EXTCODECOPY: 100,
	vm.RETURNDATASIZE: 2, vm.RETURNDATACOPY: 3, vm.BLOCKHASH: 20, vm.COINBASE: 2, vm.TIMESTAMP: 2, vm.NUMBER: 2,
	vm.DIFFICULTY: 2, vm.GASLIMIT: 2, vm.POP: 2, vm.MLOAD: 3, vm.MSTORE: 3, vm.MSTORE8: 3, vm.SLOAD: 100, vm.SSTORE: 20000,
	vm.JUMP: 8, vm.JUMPI: 10, vm.PC: 2, vm.MSIZE: 2, vm.GAS: 2, vm.JUMPDEST: 1,
	vm.PUSH1: 3, vm.PUSH2: 3, vm.PUSH3: 3, vm.PUSH4: 3, vm.PUSH5: 3, vm.PUSH6: 3, vm.PUSH7: 3, vm.PUSH8: 3, vm.PUSH9: 3,
	vm.PUSH10: 3, vm.PUSH11: 3, vm.PUSH12: 3, vm.PUSH13: 3, vm.PUSH14: 3, vm.PUSH15: 3, vm.PUSH16: 3, vm.PUSH17: 3,
	vm.PUSH18: 3, vm.PUSH19: 3, vm.PUSH20: 3, vm.PUSH21: 3, vm.PUSH22: 3, vm.PUSH23: 3, vm.PUSH24: 3, vm.PUSH25: 3,
	vm.PUSH26: 3, vm.PUSH27: 3, vm.PUSH28: 3, vm.PUSH29: 3, vm.PUSH30: 3, vm.PUSH31: 3, vm.PUSH32: 3,
	vm.DUP1: 3, vm.DUP2: 3, vm.DUP3: 3, vm.DUP4: 3, vm.DUP5: 3, vm.DUP6: 3, vm.DUP7: 3, vm.DUP8: 3, vm.DUP9: 3,
	vm.DUP10: 3, vm.DUP11: 3, vm.DUP12: 3, vm.DUP13: 3, vm.DUP14: 3, vm.DUP15: 3, vm.DUP16: 3,
	vm.SWAP1: 3, vm.SWAP2: 3, vm.SWAP3: 3, vm.SWAP4: 3, vm.SWAP5: 3, vm.SWAP6: 3, vm.SWAP7: 3, vm.SWAP8: 3,
	vm.SWAP9: 3, vm.SWAP10: 3, vm.SWAP11: 3, vm.SWAP12: 3, vm.SWAP13: 3, vm.SWAP14: 3, vm.SWAP15: 3, vm.SWAP16: 3,
	vm.LOG0: 375, vm.LOG1: 750, vm.LOG2: 1125, vm.LOG3: 1500, vm.LOG4: 1875,
	vm.CREATE: 32000, vm.CALL: 700, vm.CALLCODE: 700, vm.DELEGATECALL: 700, vm.STATICCALL: 700, vm.RETURN: 0,
	vm.REVERT: 0, vm.SELFDESTRUCT: 5000,
}

// GetGasCost returns gas for opcode, default 0 if unknown
func GetGasCost(op vm.OpCode) uint64 {
	if cost, ok := GasTable[op]; ok {
		return cost
	}
	return 0
}
