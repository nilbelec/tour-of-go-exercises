package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	current, next := 0, 1
	return func() int {
		defer func() {
			current, next = next, current+next
		}()
		return current
	}
}

// exercise04 run exercise 04
func exercise04() {
	f := fibonacci()
	for i := 0; i < 90; i++ {
		fmt.Println(f())
	}
}
