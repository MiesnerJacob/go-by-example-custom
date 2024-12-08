package main

import "fmt"

func main() {
	// Define array
	var numbers_arrray [10]int
	fmt.Println(numbers_arrray)

	// Set array value
	numbers_arrray[0] = 1
	fmt.Println(numbers_arrray)
	fmt.Println(numbers_arrray[0])

	// Using built-in length function
	fmt.Println("Length:", len(numbers_arrray))

	// Declare and initialize array in one line
	integers_array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(integers_array)

	// Infer array length
	strings_array := [...]string{"a", "b", "c"}
	fmt.Println(strings_array)

	// Specifying array index
	floats_array := [10]float64{1: 10.1, 5: 20.2, 9: 30.3}
	fmt.Println(floats_array)

	// Multidimensional array
	var two_dimensional_array [2][2]int
	fmt.Println(two_dimensional_array)

	// Define and initialize multi-dimensional array
	multi_dimensional_array := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(multi_dimensional_array)
}
