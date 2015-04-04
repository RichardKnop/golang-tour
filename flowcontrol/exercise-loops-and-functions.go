package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	delta := 0.01
	previousZ := 0.0
	z := x / 3
	for z - previousZ > delta {
		previousZ = z
		z = z - (z * z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
}