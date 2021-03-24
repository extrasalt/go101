package main

import "fmt"

func fmap(slice []int, f func(int) int) []int {
	result := []int{}
	for _, v := range slice {
		result = append(result, f(v))
	}

	return result
}

func square(n int) int {
	return n * n
}

func main() {
	o := fmap([]int{3, 5, 9}, square)

	fmt.Println(o)

}
