package main

import (
	"fmt"
	"time"
)

func main() {

	//You don't need a wait group here because the channel receive is blocking already.

	c := make(chan string, 2)

	go func(in chan string) {
		time.Sleep(10 * time.Second)
		in <- "This gets printed after 10 seconds"
	}(c)

	fmt.Println("This gets printed immediately")

	fromGoRoutine := <-c
	fmt.Println(fromGoRoutine)

}
