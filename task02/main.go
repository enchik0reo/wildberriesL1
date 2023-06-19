// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import (
	"fmt"
	"sync"
)

// Решены проблемы:
// переменной цикла в замыкании (передаём копию переменной цикла)
// синхронизации (используем WaitGroup)
func printSquare(array [5]int) {
	wg := sync.WaitGroup{}
	for _, j := range array {
		wg.Add(1)
		go func(j int) {
			fmt.Printf("%d\n", j*j)
			wg.Done()
		}(j)
	}
	wg.Wait()
}

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	printSquare(array)
}
