package main

import "fmt"

func main() {
	// Basic switch
	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	}

	// Multiple expressions
	switch i {
	case 1, 2:
		fmt.Println("One or Two")
	case 3:
		fmt.Println("Three")
	}

	// Default case
	switch i {
	case 1:
		fmt.Println("One")
	default:
		fmt.Println("Default")
	}

	// Replicating if else logic
	switch {
	case i == 1:
		fmt.Println("One")
	default:
		fmt.Println("Default")
	}

	// Type switch
	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown")
	}
}
