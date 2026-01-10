package analyzer

type StorageTracker struct {
	SLoadCount  map[uint64]int
	SStoreCount map[uint64]int
}

func NewStorageTracker() *StorageTracker {
	return &StorageTracker{
		SLoadCount:  make(map[uint64]int),
		SStoreCount: make(map[uint64]int),
	}
}
