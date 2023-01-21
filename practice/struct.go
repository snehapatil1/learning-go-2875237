package main

import "fmt"

func main() {
	poodle := Dog{"Poodle", 50}
	fmt.Println(poodle)
	fmt.Printf("%+v \n", poodle)
	fmt.Printf("Name: %v, Weight: %v \n", poodle.Name, poodle.Weight)
	poodle.Weight = 100
	fmt.Printf("Name: %v, Weight: %v \n", poodle.Name, poodle.Weight)
}

// Dog is a struct
type Dog struct {
	Name   string
	Weight int
}
