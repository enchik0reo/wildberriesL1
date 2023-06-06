// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func swapWords1(str string) string {
	runes := []rune(str)
	words := []string{}
	answ := make([]rune, 0, len(runes))
	s := []rune{}
	for i, j := range runes {
		if j == ' ' {
			words = append(words, string(s))
			s = []rune{}
		} else if i == len(runes)-1 {
			s = append(s, j)
			words = append(words, string(s))
		} else {
			s = append(s, j)
		}
	}

	for i := len(words) - 1; i >= 0; i-- {
		answ = append(answ, []rune(words[i])...)
		answ = append(answ, ' ')
	}

	return string(answ)
}

// С помощью функции strings.Split() и структуры strings.Builder{}
func swapWords2(str string) string {
	words := strings.Split(str, " ")
	b := strings.Builder{}
	b.Grow(len(str))
	for i := len(words) - 1; i >= 0; i-- {
		b.WriteString(fmt.Sprintf("%s ", words[i]))
	}

	return b.String()
}

func main() {
	str := "snow dog sun"
	fmt.Println(swapWords1(str))
	fmt.Println(swapWords2(str))
}
