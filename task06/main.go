// Реализовать все возможные способы остановки выполнения горутины.
package main

import (
	"fmt"
	"time"
)

// --------------------------------------------
// Выход из main функции
func fakePrinter() {
	fmt.Println("From fake: I'll print forewer!")
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("From fake: Yes I will")
	}
}

func main() {
	go fakePrinter()
	time.Sleep(3 * time.Second)
	fmt.Println("From main: I just stopped you, fake")
}

// --------------------------------------------
// Использование канала для остановки
/* func badPrinter(knock <-chan struct{}) {
	for {
		select {
		case <-knock:
			defer fmt.Println("done!")
			return
		default:
			fmt.Print("d ")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	knock := make(chan struct{})
	go badPrinter(knock)
	time.Sleep(3 * time.Second)
	knock <- struct{}{}
	time.Sleep(100 * time.Millisecond)
} */

// --------------------------------------------
// Использование контекста для остановки
/* func bugPrinter(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			defer fmt.Println("From printer: I'm fixed!")
			return
		default:
			fmt.Println("From printer: err")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go bugPrinter(ctx)
	time.Sleep(4 * time.Second)
	fmt.Println("From main: Everything is under control")
} */
