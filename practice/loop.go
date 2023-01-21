package main

import "fmt"

func main() {
	colors := make([]string, 3)
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"

	// Traditional for loop
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// Range for loop
	for i := range colors {
		fmt.Println(colors[i])
	}

	// Substitute for while loop
	i := 0
	for i < len(colors) {
		fmt.Println(colors[i])
		i++
	}

	// Break the loop
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum: ", sum)
		if sum > 200 {
			goto theEnd
		}
	}
theEnd:
	fmt.Println("End of program.")
}
