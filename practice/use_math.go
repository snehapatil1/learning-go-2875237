package main

import (
	"fmt"
	"math"
)

func main() {
	// Print all numbers
	f1, f2, f3 := 10.10, 20.20, 30.30
	fmt.Printf("Numbers are: %v, %v, %v\n", f1, f2, f3)

	// Sum all numbers
	fSum := f1 + f2 + f3
	fmt.Println("Sum is: ", fSum)

	// Round the sum to 2 decimals using formatter
	fmt.Printf("Sum rounded to 2 decimals: %.2f\n", fSum)

	// Round the sum to 2 decimals using math function
	rounded := math.Round(fSum)
	fmt.Println("Rounded sum using math: ", rounded)

	// Setting variable already initialised
	fSum = fSum + 100
	fmt.Println("Setting variable already initialised: ", fSum)

}
