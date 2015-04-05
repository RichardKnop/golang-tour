package main

import "fmt"

func main() {
	var z []int
	fmt.Println(len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}