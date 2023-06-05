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

	sortNum := sortNum(nums)

	quickSort(nums[:sortNum])
	quickSort(nums[sortNum:])

	return nums
}

func sortNum(nums []int) int {
	halfNum := nums[len(nums)/2]
	minIdx := 0
	maxIdx := len(nums) - 1

	for {
		if nums[minIdx] < halfNum {
			minIdx++
		}

		if nums[maxIdx] > halfNum {
			maxIdx--
		}

		if minIdx >= maxIdx {
			return maxIdx
		}

		nums[minIdx], nums[maxIdx] = nums[maxIdx], nums[minIdx]
	}
}

func main() {
	nums := []int{7, 5, 8, 4, 0, 9, 3, 1, 2}
	quickSort(nums)
	fmt.Println(nums)

	// С помощью пакета sort
	nums = []int{7, 5, 8, 4, 0, 9, 3, 1, 2}
	sort.Ints(nums)
	fmt.Println(nums)
}
