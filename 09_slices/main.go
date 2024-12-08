package main

import (
	"fmt"
	"slices"
)

func main() {
	// Define slice
	var numbers_slice []int
	fmt.Println(numbers_slice)

	// Define slice with max capacity
	numbers_slice = make([]int, 10)
	fmt.Println(numbers_slice)

	// Set and get slice values
	numbers_slice[0] = 1
	numbers_slice[1] = 2
	numbers_slice[2] = 3
	fmt.Println(numbers_slice)
	fmt.Println(numbers_slice[0])

	// Slice append
	numbers_slice = append(numbers_slice, 4, 5, 6)
	fmt.Println(numbers_slice)

	// Slice copy
	numbers_slice_copy := make([]int, len(numbers_slice))
	copy(numbers_slice_copy, numbers_slice)
	fmt.Println(numbers_slice_copy)

	// Slice a slice
	numbers_slice_slice := numbers_slice[1:4]
	fmt.Println(numbers_slice_slice)

	numbers_slice_slice = numbers_slice[4:]
	fmt.Println(numbers_slice_slice)

	numbers_slice_slice = numbers_slice[:4]
	fmt.Println(numbers_slice_slice)

	// Declare and initialize slice in one line
	string_slice := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(string_slice)

	// Slice utility functions
	fmt.Println(slices.Contains(string_slice, "a"))
	fmt.Println(slices.Index(string_slice, "a"))

	// Multidimensional slice
	multi_dimensional_slice := [][]int{{1, 2}, {3, 4}}
	fmt.Println(multi_dimensional_slice)
}
