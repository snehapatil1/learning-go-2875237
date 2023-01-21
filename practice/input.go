package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Initialize a reader and specify from where the input will be given
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text Here: ")

	// Get input from the reader and use a delimiter '\n' to specify EOF
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered text is as follows: ", input)
}
