// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	value uint64
}

func incrementWithMutex(c *counter, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	c.value++
	mu.Unlock()
	wg.Done()
}

func incrementWithAtomic(c *counter, wg *sync.WaitGroup) {
	atomic.AddUint64(&c.value, 1)
	wg.Done()
}

func main() {
	c := new(counter)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go incrementWithMutex(c, &wg, &mu)
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go incrementWithAtomic(c, &wg)
	}

	wg.Wait()

	fmt.Println(c.value)
}
