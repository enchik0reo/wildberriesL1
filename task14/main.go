// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

// С помощью switch type описать все варианты
func whatIsIt1(something interface{}) {
	switch something.(type) {
	case int:
		fmt.Println("This is int")
	case string:
		fmt.Println("This is string")
	case bool:
		fmt.Println("This is bool")
	case chan int:
		fmt.Println("This is int channel")
	// some cases ...
	default:
		fmt.Printf("I don't know this type\n")
	}
}

// С помощью reflect.TypeOf()
func whatIsIt2(something interface{}) {
	fmt.Println(reflect.TypeOf(something))
}

// С помощью fmt.Printf()
func whatIsIt3(something interface{}) {
	fmt.Printf("%T\n", something)
}

func main() {
	whatIsIt1(make(chan int))
	whatIsIt2(make(chan int))
	whatIsIt3(make(chan int))
}
