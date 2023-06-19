// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2....) с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
)

func sumSquare1(slice []int) int {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	sum := 0
	for _, j := range slice {
		wg.Add(1)
		go func(j int) {
			mu.Lock()
			sum += j * j
			mu.Unlock()
			wg.Done()
		}(j)
	}
	wg.Wait()

	return sum
}

func sumSquare2(slice []int) int {
	ch := make(chan int)
	sum := 0

	for _, j := range slice {
		go func(j int, ch chan<- int) {
			ch <- j * j
		}(j, ch)
	}

	for i := 0; i < len(slice); i++ {
		sum += <-ch
	}

	return sum
}

func main() {
	slice := []int{2, 4, 6, 8, 10}
	fmt.Println(sumSquare1(slice))
	fmt.Println(sumSquare2(slice))
}
