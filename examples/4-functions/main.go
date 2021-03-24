package main

import (
	"errors"
	"fmt"
)

//can also be written as (a, b int)
func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return a / b, nil
}

func addALot(nums ...int) int {
	sum := 0
	for _, a := range nums {
		sum += a
	}

	return sum
}

func recursive(n int) int {
	if n == 0 {
		return 0
	}
	m := n - 1
	return n + recursive(m)
}

func main() {
	fmt.Println(add(3, 4))

	fmt.Println(addALot(1, 2, 2, 2, 2))

	//Multiple returns
	q, err := divide(9, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(q)

	//Recursive
	fmt.Println(recursive(10))
}
