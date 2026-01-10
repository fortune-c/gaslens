package analyzer

type Loop struct {
	StartPC int
	EndPC   int
	Count   int
}

type LoopTracker struct {
	Loops []Loop
}

func NewLoopTracker() *LoopTracker {
	return &LoopTracker{
		Loops: []Loop{},
	}
}

func (lt *LoopTracker) RecordLoop(start, end int) {
	for i := range lt.Loops {
		if lt.Loops[i].StartPC == start && lt.Loops[i].EndPC == end {
			lt.Loops[i].Count++
			return
		}
	}
	lt.Loops = append(lt.Loops, Loop{
		StartPC: start,
		EndPC:   end,
		Count:   1,
	})
}
