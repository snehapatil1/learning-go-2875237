package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Initialise reader input
	reader := bufio.NewReader(os.Stdin)

	// Input value 1 as string and convert to float
	fmt.Print("Enter Value 1: ")
	val1, _ := reader.ReadString('\n')
	floatval1, err := strconv.ParseFloat(strings.TrimSpace(val1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be an integer/float value.")
	}

	// Input value 2 as string and convert to float
	fmt.Print("Enter Value 2: ")
	val2, _ := reader.ReadString('\n')
	floatval2, err := strconv.ParseFloat(strings.TrimSpace(val2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be an integer/float value.")
	}

	// Sum of value 1 and value 2 and round the sum
	fsum := floatval1 + floatval2
	sum := math.Round(float64(fsum))
	fmt.Printf("Sum of %v and %v is: %v", floatval1, floatval2, sum)
}
