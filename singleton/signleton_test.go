package singleton

import (
	"sync"
	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)

	instances := [parCount]*Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
