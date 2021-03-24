package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Hello world %d\n", i)
		}
	}

	j := 8
	for j < 10 {
		fmt.Printf("Like a while loop %d\n", j)
		j++
	}

	aSlice := []int{1, 5, 7, 9}

	for i, v := range aSlice {
		fmt.Printf("%d-%d\n", i, v)
	}

}
