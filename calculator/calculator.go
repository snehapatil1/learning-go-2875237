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
	// fsum := floatval1 + floatval2
	// sum := math.Round(float64(fsum))
	// fmt.Printf("Sum of %v and %v is: %v", floatval1, floatval2, sum)

	// Input value 3 as string to select operation from (+, -, *, /)
	fmt.Print("Select operation from(+, -, *, /): ")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	op := 0.00
	switch operation {
	case "+":
		op = addition(floatval1, floatval2)
		fmt.Printf("Addition of %v and %v is %v", floatval1, floatval2, math.Round(float64(op)))
	case "-":
		op = subtraction(floatval1, floatval2)
		fmt.Printf("Subtraction of %v and %v is %v", floatval1, floatval2, math.Round(float64(op)))
	case "*":
		op = multiplication(floatval1, floatval2)
		fmt.Printf("Multiplication of %v and %v is %v", floatval1, floatval2, math.Round(float64(op)))
	case "/":
		op = division(floatval1, floatval2)
		fmt.Printf("Division of %v and %v is %v", floatval1, floatval2, math.Round(float64(op)))
	default:
		panic("This is not a valid operation! Please select one of (+, -, *, /).")
	}

}

func addition(value1 float64, value2 float64) float64 {
	return value1 + value2
}

func subtraction(value1 float64, value2 float64) float64 {
	return value1 - value2
}

func multiplication(value1 float64, value2 float64) float64 {
	return value1 * value2
}

func division(value1 float64, value2 float64) float64 {
	return value1 / value2
}
