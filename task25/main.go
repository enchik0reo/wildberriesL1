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
	d := 2 * time.Second
	fmt.Println("Start sleep")
	sleep(d)
	fmt.Println("Wake up after", d)
}
