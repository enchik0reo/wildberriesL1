// Удалить i-ый элемент из слайса.
package main

import "fmt"

// С созданием нового базового массива, изменяет порядок в слайсе
func del1[T any](slice []T, i uint) []T {
	s := make([]T, len(slice))
	copy(s, slice)
	s[i] = s[len(s)-1]

	return s[:len(s)-1]
}

// С созданием нового базового массива, не изменяет порядок в слайсе
func del2[T any](slice []T, i uint) []T {
	sl1 := slice[:i]
	sl2 := slice[i+1:]
	s := make([]T, 0, len(slice)-1)
	s = append(s, sl1...)
	s = append(s, sl2...)

	return s
}

// Без создания нового базового массива
func del3[T any](slice []T, i uint) []T {
	copy(slice[i:], slice[i+1:])

	return slice[:len(slice)-1]
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(del1(slice, 2))
	fmt.Println(del2(slice, 2))
	fmt.Println("Начальный массив:", slice)

	fmt.Println(del3(slice, 2))
	fmt.Println("Начальный массив:", slice)
}
