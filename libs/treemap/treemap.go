package treemap

import (
	"sync"
)

// TreeMap tree map
type TreeMap struct {
	OrderMap *Tree
	mutex    sync.RWMutex
}

// NewMap return a pointer to tree map
func NewMap() *TreeMap {
	tm := &TreeMap{
		OrderMap: NewTree(),
	}
	return tm
}

// Store store the key-value
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

// Delete delete the key-value
func (tm *TreeMap) Delete(key Keytype) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	tm.OrderMap.Delete(key)
}

// Empty empty or not
func (tm *TreeMap) Empty() bool {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	return tm.OrderMap.Empty()
}

// Min return the top
func (tm *TreeMap) Min() *node {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	return tm.OrderMap.Iterator()
}

// Size return the size of tree map
func (tm *TreeMap) Size() int {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	return tm.OrderMap.Size()
}

// Clear clear the tree map
func (tm *TreeMap) Clear() {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	tm.OrderMap.Clear()
}
