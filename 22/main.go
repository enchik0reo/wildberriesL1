// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 2^20 = 1048576
	a := new(big.Int)
	b := new(big.Int)

	a.SetUint64(60485766048576)
	b.SetUint64(20485762048576)

	fmt.Printf("num1 = %v; num2 = %v\n", a, b)

	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Printf("Multi = %v\n", mul)

	div := new(big.Int)
	div.Div(a, b)
	fmt.Printf("Div = %v\n", div)

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("Sum = %v\n", sum)

	dif := new(big.Int)
	dif.Sub(a, b)
	fmt.Printf("Dif = %v\n", dif)
}
