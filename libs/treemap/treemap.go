package treemap

import (
	"sync"
)

type TreeMap struct {
	OrderMap *Tree
	mutex    sync.RWMutex
}

func NewMap() *TreeMap {
	tm := &TreeMap{
		OrderMap: NewTree(),
	}
	return tm
}

func (tm *TreeMap) Store(key Keytype, value interface{}) bool {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	update := false
	if tm.OrderMap.FindIter(key) != nil {
		tm.OrderMap.Delete(key)
		update = true
	}
	tm.OrderMap.Insert(key, value)
	return update
}

func (tm *TreeMap) Delete(key Keytype) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	tm.OrderMap.Delete(key)
}

func (tm *TreeMap) Empty() bool {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	return tm.OrderMap.Empty()
}

func (tm *TreeMap) Min() *node {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	return tm.OrderMap.Iterator()
}

func (tm *TreeMap) Size() int {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	return tm.OrderMap.Size()
}

func (tm *TreeMap) Clear() {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	tm.OrderMap.Clear()
}
