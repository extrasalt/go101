package main

import "fmt"

func main() {
	aMap := map[string]int{"one": 1, "two": 2}

	aMap["three"] = 3

	// aMap[4] = "four" // type-safe

	for k, v := range aMap {
		fmt.Printf("%s, %v\n", k, v)
	}

}
