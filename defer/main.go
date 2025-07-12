package main

import "fmt"
  
func main() {
	i := 0
	// Key Principle: Deferred Arguments Are Evaluated Immediately
	// The moment defer fmt.Println(i) runs, Go captures the current value of i (which is 0) and stores it.
	// It doesn’t wait for the end of the function to read i—it already has the value.
	// Later, i++ increments the variable, but that doesn’t change what the deferred call stored.
	// When the function returns, Go executes the deferred call and prints the value it captured: 0
	defer fmt.Println(i)
	i++
	// Captures 1
	defer_func()
	return
}

func defer_func() {
	i := 0
	defer func() {
		fmt.Println(i)
	}()
	i++
	return
}