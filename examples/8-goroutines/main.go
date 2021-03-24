package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//You don't need a wait group here because the channel receive is blocking already.
	var wg sync.WaitGroup

	c := make(chan string, 2)

	wg.Add(1)

	go func(in chan string) {
		defer wg.Done()
		time.Sleep(10 * time.Second)
		in <- "This gets printed after 10 seconds"
	}(c)

	fmt.Println("This gets printed immediately")

	fromGoRoutine := <-c
	fmt.Println(fromGoRoutine)

	wg.Wait()

}
