package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("x is of type %T\n", x)

	y := 3.123
	fmt.Printf("y is of type %T\n", y)

	z := 5 + 8i
	fmt.Printf("z is of type %T\n", z)
}