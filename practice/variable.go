package main

import (
	"fmt"
)

const myInt2 = 10

func main() {

	// Method 1: Print String variable value and type
	var myString1 = "Sneha"
	fmt.Println(myString1)
	fmt.Printf("Type of variable is %T\n", myString1)

	// Method 2: Print String variable value and type
	myString2 := "Sneha Patil"
	fmt.Println(myString2)
	fmt.Printf("Type of variable is %T\n", myString2)

	// Print Int variable value and type
	myInt1 := 100
	fmt.Println(myInt1)
	fmt.Printf("Type of variable is %T\n", myInt1)

	// Print Constant Int variable value and type
	fmt.Println(myInt2)
	fmt.Printf("Type of const variable is %T\n", myInt2)

}
