package main

import (
	"fmt"
	"math"
)

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

func power(p int) func(int) int {
	return func(n int) int {
		return int(math.Pow(float64(n), float64(p)))
	}
}

func main() {
	o := fmap([]int{3, 5, 9}, square)
	fmt.Println(o)

	cube := power(3)
	o = fmap([]int{3, 5, 9}, cube)
	fmt.Println(o)

}
