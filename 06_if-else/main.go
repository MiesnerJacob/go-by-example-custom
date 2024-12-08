package main

import "fmt"

func main() {
	// If statement
	if dog := "Coco"; dog == "Coco" {
		fmt.Println("Coco is a pug")
	}

	// If else statement
	if dog := "Coco"; dog == "Coco" {
		fmt.Println("Hello Coco")
	} else {
		fmt.Println("Coco is not here")
	}

	// If else if statement
	if dog := "Coco"; dog == "Pug" {
		fmt.Println("Hello Coco")
	} else if dog == "Cleo" {
		fmt.Println("Cleo, where is Coco?")
	} else {
		fmt.Println("Who is this?")
	}
}
