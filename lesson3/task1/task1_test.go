package task1

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {

	runtime.GOMAXPROCS(10)

	concurrentMap := NewSafeMap()

	var wg sync.WaitGroup

	wg.Add(2000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			concurrentMap.Set(fmt.Sprintf("%d", rand.Intn(100)), rand.Intn(100))
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			concurrentMap.Get(fmt.Sprintf("%d", rand.Intn(100)))
		}()
	}

	wg.Wait()

	var wg2 sync.WaitGroup

	wg2.Add(100)

	for i := 0; i < 100; i++ {
		concurrentMap.Set(fmt.Sprintf("%d", i), i)
	}

	for i := 0; i < 100; i++ {
		i := i
		go func() {
			defer wg2.Done()
			val := concurrentMap.Get(fmt.Sprintf("%d", i))

			if val.(int) != i {
				t.Errorf("Expected %d, got %d", i, val)
			}
		}()
	}
	wg2.Wait()
}
