// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.
package main

import (
	"fmt"
	"sort"
)

// O = n*log2n в среднем, O = n^2 в худшем
func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	var less, greater []int
	for _, num := range nums[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	res := append(quickSort(less), pivot)
	res = append(res, quickSort(greater)...)
	return res
}

func main() {
	nums := []int{7, 5, 8, 4, 0, 9, 3, 1, 2}
	fmt.Println(quickSort(nums))

	// С помощью пакета sort
	sort.Ints(nums)
	fmt.Println(nums)
}
