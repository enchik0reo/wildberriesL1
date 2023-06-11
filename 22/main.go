// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 2^20 = 1048576
	x := new(big.Int)
	y := new(big.Int)
	ansv := new(big.Int)

	x.SetString("10000000000000000000", 10)
	y.SetString("5000000000000000000", 10)
	fmt.Printf("num1 = %v; num2 = %v\n", x, y)

	ansv.Mul(x, y)
	fmt.Printf("Multi = %v\n", ansv)

	ansv.Div(x, y)
	fmt.Printf("Div = %v\n", ansv)

	ansv.Add(x, y)
	fmt.Printf("Sum = %v\n", ansv)

	ansv.Sub(x, y)
	fmt.Printf("Dif = %v\n", ansv)
}
