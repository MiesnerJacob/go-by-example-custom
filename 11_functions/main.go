package main

import "fmt"

// Define function
func hello_world() {
	fmt.Println("Hello, World!")
}

// Define function with return value
func hello(name string) string {
	return "Hello, " + name + "!"
}

// Define function with multiple parameters of same type
func add_integers(a, b int) int {
	return a + b
}

// Run function
func main() {
	hello_world()

	hello_response := hello("John")
	fmt.Println(hello_response)

	integer_sum := add_integers(1, 2)
	fmt.Println(integer_sum)
}
