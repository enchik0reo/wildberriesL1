// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

// либо использовать структуру sync.Map
func save(s string, i int, m map[int]string, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	m[i] = s
	mu.Unlock()
	wg.Done()
}

func main() {
	m := make(map[int]string)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go save("yup", i, m, &wg, &mu)
	}
	wg.Wait()
	fmt.Println(len(m))
}
