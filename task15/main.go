// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

/* var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
} */

// 1. justString скорее всего в куче, дорого будет очищать после неё память, доступна в любом месте кода
// 2. Строки не изменяемы, зачем то пытаемся это делать, более явно создать копию
// 3. Скорее всего хотим получить строку длиной в 100 рун, а не байт, поэтому преобразуем в []rune
package main

import "fmt"

func createHugeString(n int) string {
	str := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		str = append(str, 'ы')
	}
	return string(str)
}

func someFunc() string {
	str := createHugeString(1 << 10)
	runs := []rune(str)[:100]
	return string(runs)
}

func main() {
	justString := someFunc()

	fmt.Println(justString)
}
