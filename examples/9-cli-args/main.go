package main

import (
	"fmt"
	"os"
	"strconv"
)

func square(n int) int {
	return n * n
}

func cube(n int) int {
	return n * n * n
}

func main() {

	nStr := os.Args[2]
	n, _ := strconv.Atoi(nStr)

	switch os.Args[1] {
	case "square":
		fmt.Println(square(n))
	case "cube":
		fmt.Println(cube(n))
	default:
		fmt.Println(n)
	}
}
