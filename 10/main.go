// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
package main

import "fmt"

func main() {
	t := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5.3, 0.2, -0.7, -9.9, 102.2, 0}
	groups := make(map[int][]float64)

	for _, temp := range t {
		step := int(temp/10) * 10
		if step == 0 {
			if temp > 0 {
				step = 1
			} else if temp < 0 {
				step = -1
			}
		}

		groups[step] = append(groups[step], temp)
	}

	for step, temp := range groups {
		fmt.Printf("%d: %v\n", step, temp)
	}
}
