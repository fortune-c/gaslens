package analyzer

type FunctionInfo struct {
	Selector string
	EntryPC  int
	Gas      uint64
}

type FunctionTracker struct {
	Functions []FunctionInfo
}

func NewFunctionTracker() *FunctionTracker {
	return &FunctionTracker{
		Functions: []FunctionInfo{},
	}
}

func (ft *FunctionTracker) AddFunction(selector string, pc int) {
	ft.Functions = append(ft.Functions, FunctionInfo{
		Selector: selector,
		EntryPC:  pc,
		Gas:      0,
	})
}

func (ft *FunctionTracker) AddGas(pc int, gas uint64) {
	for i := range ft.Functions {
		if pc >= ft.Functions[i].EntryPC {
			ft.Functions[i].Gas += gas
		}
	}
}
