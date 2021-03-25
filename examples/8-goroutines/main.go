package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Second)
		fmt.Println("This gets printed after 10 seconds")
	}()

	fmt.Println("This gets printed immediately")

	wg.Wait()

}
