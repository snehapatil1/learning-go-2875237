package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	output, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println("ERROR! : ", err)
	} else {
		fmt.Println("You entered: ", output)
		fmt.Printf("Type of output is %T\n", output)
	}
}
