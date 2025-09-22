package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type CustomSyncMap interface {
	Get(key string) (int, bool)
	Set(key string, value int)
	Delete(key string)
	Len() int
}

type mapWrapper struct {
	mp map[string]int
	mu sync.RWMutex
}

func NewSyncMap() CustomSyncMap {
	return &mapWrapper{
		mp: make(map[string]int),
	}
}

func (m *mapWrapper) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	val, ok := m.mp[key]
	return val, ok
}

func (m *mapWrapper) Set(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.mp[key] = value
}

func (m *mapWrapper) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.mp, key)
}

func (m *mapWrapper) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.mp)
}

func main() {
	sm := NewSyncMap()
	var wg sync.WaitGroup

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				k := fmt.Sprintf("k-%d", rand.Intn(1000))
				sm.Set(k, j+id)
				if j%7 == 0 {
					sm.Delete(k)
				}
			}
		}(i)
	}

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				k := fmt.Sprintf("k-%d", rand.Intn(1000))
				_, _ = sm.Get(k)
				_ = sm.Len()
			}
		}()
	}

	wg.Wait()
}
