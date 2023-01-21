package main

import "fmt"

func main() {
	function1()
	sum := function2(10, 20)
	fmt.Println("Sum from Function 2: ", sum)
	multisum := function3(1, 2, 3, 4, 5)
	fmt.Println("Sum from Function 3: ", multisum)
	multisum1, val_len := function4(1, 2, 3, 4, 5, 6)
	fmt.Println("Sum from Function 4: ", multisum1, val_len)
}

func function1() {
	fmt.Println("This is Function 1.")
}

func function2(value1 int, value2 int) int {
	return value1 + value2
}

func function3(value ...int) int {
	sum := 0
	for _, v := range value {
		sum += v
	}
	return sum
}

func function4(value ...int) (int, int) {
	sum := 0
	for _, v := range value {
		sum += v
	}
	return sum, len(value)
}
