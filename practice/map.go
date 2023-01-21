package main

import (
	"fmt"
)

func main() {
	hashmap := make(map[string]string)
	hashmap["R"] = "Red"
	hashmap["B"] = "Blue"
	hashmap["G"] = "Green"

	fmt.Println(hashmap)

	hashmap["P"] = "Pink"
	fmt.Println(hashmap)

	delete(hashmap, "P")
	fmt.Println(hashmap)

	for k, v := range hashmap {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(hashmap))
	i := 0
	for k := range hashmap {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(hashmap[keys[i]])
	}
}
