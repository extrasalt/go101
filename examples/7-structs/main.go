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
	engine Engine
}

func main() {

	car := Car{
		engine: Engine{
			cc: 1500,
		},
	}

	car.engine.start()

	fmt.Println(car.engine.state)

}
