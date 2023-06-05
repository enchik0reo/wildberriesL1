// Дана переменная int64. Разработать программу которая устанавливает i-й бит в
// 1 или 0.
package main

import (
	"fmt"
)

func changeBit(n int64) int64 {
	var i uint8
	fmt.Print("Бит №: ")
	fmt.Scan(&i)
	if i > 63 {
		panic("Бит от 0 до 63")
	}
	var want uint8
	fmt.Println("1 или 0: ")
	fmt.Scan(&want)
	if want > 1 {
		panic("1 или 0")
	}

	mask := 1 << i
	if want == 1 {
		// побитовое OR
		n |= int64(mask)
	} else {
		// побитовое AND XOR
		n &^= int64(mask)
	}

	return n
}

func main() {
	var n int64 = 240
	fmt.Printf("Было: %b, стало: %b\n", n, changeBit(n))
}
