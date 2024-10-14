package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

type ConcurrentMap struct {
	mu sync.Mutex
	mp map[int]int
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		mp: make(map[int]int),
	}
}

func (m *ConcurrentMap) writeToMap(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp[i] = i * i
}

func writeToSyncMap(m *sync.Map, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Store(i, i*i)
}

func main() {
	var wg sync.WaitGroup

	syncMap := new(sync.Map)
	mMutex := NewConcurrentMap()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go mMutex.writeToMap(i, &wg)
		wg.Add(1)
		go writeToSyncMap(syncMap, i, &wg)
	}

	wg.Wait()

	fmt.Println("Contents of sync.Map:")
	syncMap.Range(func(key, value any) bool {
		fmt.Printf("syncMap[%d] = %d\n", key, value)
		return true
	})

	fmt.Println("\nContents of map with Mutex:")
	for key, value := range mMutex.mp {
		fmt.Printf("m[%d] = %d\n", key, value)
	}
}
