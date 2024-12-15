package main

import "fmt"

func average(nums ...float64) float64 {
	fmt.Println(nums)

	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

func main() {
	fmt.Println(average(1, 2, 3, 4, 5))
}
