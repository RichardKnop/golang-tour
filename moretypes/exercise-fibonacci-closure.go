package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	started := false
	var last, prevLast int
	return func() int {
		if started == false {
			started = true
			last = 0
			return 0
		}
		if last == 0 {
			last = 1
			prevLast = 0
			return 1
		}

		next := last + prevLast
		prevLast = last
		last = next
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}