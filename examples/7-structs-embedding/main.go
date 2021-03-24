package main

import "fmt"

type Engine struct {
	cc    int
	state string
}

func (e *Engine) start() {
	e.state = "started"
}

type Car struct {
	Engine
}

func main() {

	car := Car{
		Engine{
			cc: 1500,
		},
	}

	car.Engine.start()
	// car.start()

	fmt.Println(car.Engine.state)
	// fmt.Println(car.state)

}
