package utility

import "sync"

// SafeMap todo 升级版本后替换为 sync.Map
type SafeMap struct {
	sync.RWMutex
	data map[string]interface{}
}

func NewSafeMap(size int) *SafeMap {
	sm := new(SafeMap)
	sm.data = make(map[string]interface{}, size)
	return sm
}

func (sm *SafeMap) Load(key string) interface{} {
	sm.RLock()
	value := sm.data[key]
	sm.RUnlock()
	return value
}

func (sm *SafeMap) Store(key string, value interface{}) {
	sm.Lock()
	sm.data[key] = value
	sm.Unlock()
}

func (sm *SafeMap) Keys() []string {
	sm.RLock()
	value := make([]string, 0)
	for k := range sm.data {
		value = append(value, k)
	}
	sm.RUnlock()
	return value
}

func (sm *SafeMap) ForEach(f func(key string, value interface{})) {
	sm.Lock()
	for s, i := range sm.data {
		f(s, i)
	}
	sm.Unlock()
}

func (sm *SafeMap) IsEmpty() bool {
	sm.RLock()
	value := false
	if len(sm.data) == 0 {
		value = true
	}
	sm.RUnlock()
	return value
}
