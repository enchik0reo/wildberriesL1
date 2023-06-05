// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.
package main

import "fmt"

func sharikov1(str string) string {
	runes := []rune(str)
	answ := make([]rune, 0, len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		answ = append(answ, runes[i])
	}
	return string(answ)
}

func sharikov2(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "главрыба"
	fmt.Println(sharikov1(str))
	fmt.Println(sharikov2(str))
}
