package main

import "fmt"

func main() {
	var arr1 [2]string
	arr1[0] = "Sneha"
	arr1[1] = "Patil"

	fmt.Println(arr1)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("Length of numbers: ", len(numbers))
}
