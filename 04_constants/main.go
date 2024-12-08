package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare a constant
	const dog_competition_winner string = "Vito the pug"
	fmt.Println(dog_competition_winner)

	// Declare a constant with inferred type
	const pi = 3.14
	fmt.Println(pi)
	fmt.Println(math.Sin(pi))

	// Declare multiple constants
	const (
		dog = "Coco"
		cat = "Tabby"
	)
	fmt.Println(dog, cat)
}
