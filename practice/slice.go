package main

import (
	"fmt"
	"sort"
)

func main() {
	// Array with no size is Slice
	var colors = []string{"Red", "Blue", "Green"}
	fmt.Println("Colors: ", colors)

	// Use append method to add values to slice
	colors = append(colors, "Pink")
	fmt.Println("Colors: ", colors)

	// Remove elements from slice using traditional array mechanism
	colors = append(colors[:len(colors)-1])
	fmt.Println("Colors: ", colors)

	numbers := make([]int, 3)
	numbers[0] = 56
	numbers[1] = 100
	numbers[2] = 23

	fmt.Println("Numbers: ", numbers)
	sort.Ints(numbers)
	fmt.Println("Numbers after sorting: ", numbers)
}
