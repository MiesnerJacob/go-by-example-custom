package main

import (
	"fmt"
	"maps"
)

func main() {
	// Create empty map
	var empty_map map[string]int
	fmt.Println(empty_map)

	// Set map values
	animal_map := map[string]string{}
	animal_map["dog"] = "bark"
	animal_map["cat"] = "meow"
	fmt.Println(animal_map)

	// Get map values
	fmt.Println(animal_map["dog"])

	// Try getting non-existent value
	fmt.Println(animal_map["bird"])

	// Map length
	fmt.Println(len(animal_map))

	// Delete map value
	delete(animal_map, "cat")
	fmt.Println(animal_map)

	// Clear map
	clear(animal_map)
	fmt.Println(animal_map)

	// Optional second return when getting a value
	value, exists := animal_map["dog"]
	fmt.Println(value, exists)

	// Declare and initialize map in one line
	animal_map = map[string]string{"dog": "bark", "cat": "meow"}
	fmt.Println(animal_map)

	// Map utility functions
	fmt.Println(maps.Equal(animal_map, animal_map))
}
