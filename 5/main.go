// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.
package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int) {
	i := 1
	for {
		ch <- i
		i++
		time.Sleep(1 * time.Second)
	}
}

func reader(ch <-chan int) {
	for {
		fmt.Printf("msg №%d\n", <-ch)
	}
}

func main() {
	fmt.Print("Время работы программы(сек): ")
	var n time.Duration
	fmt.Scan(&n)
	var ch = make(chan int)
	go sender(ch)
	go reader(ch)
	time.Sleep(n * time.Second)
}
