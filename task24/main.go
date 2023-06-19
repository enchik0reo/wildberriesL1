// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

// c^2 = a^2 + b^2
func Distance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2))
}

func main() {
	p1 := NewPoint(0, 4)
	p2 := NewPoint(3, 0)
	dist := Distance(p1, p2)
	fmt.Printf("%.1f\n", dist)
}
