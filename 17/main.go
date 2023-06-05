// Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"fmt"
	"sort"
)

// O = log2n
func binarySearch(nums []int, want int) (int, error) {
	minIdx := 0
	maxIdx := len(nums) - 1

	for minIdx <= maxIdx {
		midIdx := (minIdx + maxIdx) / 2

		if nums[midIdx] == want {
			return midIdx, nil
		}

		if nums[midIdx] > want {
			maxIdx = midIdx - 1
		} else {
			minIdx = midIdx + 1
		}
	}

	return 0, fmt.Errorf("number %d doesn't exist in the slice", want)
}

func main() {
	size := 1000
	want := 550
	evenNums := make([]int, 0, size/2)

	for i := 1; i <= size; i++ {
		if i%2 == 0 {
			evenNums = append(evenNums, i)
		}
	}

	index, err := binarySearch(evenNums, want)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("index of number %d = %d\n", want, index)
	}

	// С помощью пакета sort
	fmt.Println("with sort package:", sort.SearchInts(evenNums, want))
}
