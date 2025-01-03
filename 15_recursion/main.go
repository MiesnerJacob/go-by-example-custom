package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Print first 10 numbers in Fibonacci sequence
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}
