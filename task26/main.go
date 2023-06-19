// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.

// Например:
//
//	abcd — true
//	abCdefAaf — false
//	aabcd — false
package main

import (
	"fmt"
	"strings"
)

func unique(s string) bool {
	if len(s) == 0 {
		return false
	}
	ls := strings.ToLower(s)
	mp := make(map[rune]uint8, len([]rune(s)))
	for _, r := range ls {
		mp[r]++
		if mp[r] > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(unique("abcd"))
	fmt.Println(unique("abCdefAaf"))
	fmt.Println(unique("aAbcd"))
	fmt.Println(unique(""))
}
