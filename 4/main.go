// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// sender ждет сигнала о завершении
// когда он его получает, то закрывает канал и завершается,
// затем воркеры видят что канал закрыт и тоже завершаются
func start(n int) {
	wg := sync.WaitGroup{}
	ch := make(chan string, n)
	block := make(chan struct{})

	go sender(ch, block)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ch, i, &wg)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	block <- struct{}{}
	wg.Wait()
}

func sender(ch chan<- string, block chan struct{}) {
	for {
		select {
		case <-block:
			close(ch)
			defer log.Println("sender finished")
			return
		default:
			ch <- "Hi form the channel"
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func worker(ch <-chan string, i int, wg *sync.WaitGroup) {
	for {
		msg, ok := <-ch
		if !ok {
			defer wg.Done()
			defer log.Printf("worker №%d finished\n", i)
			return
		}
		fmt.Printf("%s; worker №%d\n", msg, i)
	}
}

func main() {
	n := 0
	fmt.Print("Количество воркеров: ")
	fmt.Scan(&n)
	if n < 1 {
		panic("Количество воркеров не может быть <= 0")
	}
	start(n)
}
