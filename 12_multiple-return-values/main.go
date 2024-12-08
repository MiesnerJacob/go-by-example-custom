package main

import "fmt"

// Define function with multiple return values
func add_subtract(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	var int1, int2 int = 1, 2
	fmt.Print("Integers: ", int1, ", ", int2, "\n")

	sum, difference := add_subtract(int1, int2)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}
