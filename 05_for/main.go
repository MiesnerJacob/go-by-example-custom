package main

import "fmt"

func main() {
	// single condition
	value := 0
	for value < 100 {
		fmt.Println(value)
		value++
		value *= 2
	}

	// initalize, condition, increment
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// loop n times
	for i := range 10 {
		fmt.Println(i, "is my favorite number")
	}

	// loop through a collection
	numbers := []int{1, 2, 3, 4, 5}
	for i, value := range numbers {
		fmt.Println(i, ":", value, "is my favorite number")
	}

	// infinite loop (with break)
	for {
		fmt.Println("infinite loop")
		break
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
