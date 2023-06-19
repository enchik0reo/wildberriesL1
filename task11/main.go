// Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

// O = 2n, есть доп затраты по памяти в виде map
func withMap[T comparable](a []T, b []T) []T {
	c := []T{}
	if len(a) > 0 && len(b) > 0 {
		m := make(map[T]struct{}, len(a))
		for _, j := range a {
			m[j] = struct{}{}
		}

		for _, j := range b {
			if _, ok := m[j]; ok {
				c = append(c, j)
			}
		}
	}
	return c
}

// O = n^2, нет доп затрат по памяти
func doubleFor[T comparable](a []T, b []T) []T {
	c := []T{}
	if len(a) > 0 && len(b) > 0 {
		for i := range a {
			for j := range b {
				if a[i] == b[j] {
					c = append(c, a[i])
				}
			}
		}
	}
	return c
}

func main() {
	a := []int{4, 1, 0, 6, 5, 8}
	b := []int{5, 8, 3, 2, 4, 7, 9}
	fmt.Println(withMap(a, b))
	fmt.Println(doubleFor(a, b))

	c := []float64{1.3, 12, 5, 33.2}
	d := []float64{1.4, 33.2, 45, 8.7, 31, 7}
	fmt.Println(withMap(c, d))
	fmt.Println(doubleFor(c, d))

	e := []string{"llama", "red", "pajama", "is", "alone", "without", "his", "mama"}
	f := []string{"baby", "llama", "hums", "a", "tune", "mama", "says", "that", "she'll", "be", "soon"}
	fmt.Println(withMap(e, f))
	fmt.Println(doubleFor(e, f))
}
