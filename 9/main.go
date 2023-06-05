// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	nums := []int{1, 2, 3, 4, 5}

	go func() {
		for _, num := range nums {
			ch1 <- num
		}
		close(ch1)
	}()

	go func() {
		for {
			if num, ok := <-ch1; ok {
				ch2 <- num * 2
			} else {
				close(ch2)
				return
			}
		}
	}()

	for {
		if num, ok := <-ch2; ok {
			fmt.Println(num)
		} else {
			return
		}
	}
}
