package main

import "fmt"

func main() {
	// declare one variable
	var pug string
	fmt.Println(pug)

	// declare variable then value
	var pug_ranking int
	pug_ranking = 1
	fmt.Println(pug_ranking)

	//declare variable and value at the same time
	var is_pug_cat bool = false
	fmt.Println(is_pug_cat)

	// declare multiple variables and values
	var pet1, pet2 string = "Turtle", "Hamster"
	fmt.Println(pet1, pet2)

	// declare variable with type inference
	dog_competition_winner := "Vito the pug"
	fmt.Println(dog_competition_winner)
}

