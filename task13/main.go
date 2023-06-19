// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

type numConstraint interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Классический swap
func swap1[T numConstraint](x, y T) (T, T) {
	// return y, x
	x, y = y, x
	return x, y
}

// Арифметически
func swap2[T numConstraint](x, y T) (T, T) {
	x += y    // x *= y
	y = x - y // y = x / y
	x -= y    // x /= y
	return x, y
}

// С помощью XOR (исключающее или)
func swap3[T numConstraint](x, y T) (T, T) {
	//        x = 01 y = 10
	x ^= y // x = 11
	y ^= x // y = 01
	x ^= y // x = 10
	return x, y
}

func main() {
	fmt.Println(swap1(1, 2))

	fmt.Println(swap2(1, 2))

	fmt.Println(swap3(1, 2))
}
