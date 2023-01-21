package main

import "fmt"

func main() {
	myNum := 100
	var p = &myNum

	fmt.Println("Value 1: ", myNum)
	fmt.Println("Pointer for Value 1: ", *p)

	*p = *p / 2
	fmt.Println("After division")
	fmt.Println("Value 1: ", myNum)
	fmt.Println("Pointer for Value 1: ", *p)

}
