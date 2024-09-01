package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(2, 3, add)
	fmt.Println(x)

	y := doMath(5, 2, subtract)
	fmt.Println(y)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func doMath(a int, b int, f func(a, b int) int) int {
	return f(a, b)
}
