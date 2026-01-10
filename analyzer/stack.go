package analyzer

type StackEngine struct {
	Stack []uint64
}

func (se *StackEngine) Push(v uint64) {
	se.Stack = append(se.Stack, v)
}

func (se *StackEngine) Pop() uint64 {
	if len(se.Stack) == 0 {
		return 0
	}
	v := se.Stack[len(se.Stack)-1]
	se.Stack = se.Stack[:len(se.Stack)-1]
	return v
}

func (se *StackEngine) Peek(n int) uint64 {
	idx := len(se.Stack) - 1 - n
	if idx < 0 {
		return 0
	}
	return se.Stack[idx]
}

func (se *StackEngine) Dup(n int) {
	v := se.Peek(n - 1)
	se.Push(v)
}

func (se *StackEngine) Swap(n int) {
	top := len(se.Stack) - 1
	idx := top - n
	if idx < 0 {
		return
	}
	se.Stack[top], se.Stack[idx] = se.Stack[idx], se.Stack[top]
}
