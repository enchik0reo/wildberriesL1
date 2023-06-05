// Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	t := 2 * time.Second
	fmt.Println("Start sleep")
	sleep(t)
	fmt.Println("Wake up after", t)
}
