// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.
package main

import "fmt"

func set(s []string) []string {
	if len(s) < 2 {
		return s
	}

	m := make(map[string]struct{}, len(s))
	set := []string{}

	for _, j := range s {
		if _, ok := m[j]; !ok {
			m[j] = struct{}{}
			set = append(set, j)
		}
	}

	return set
}

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(set(slice))
}
