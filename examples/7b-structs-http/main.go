package main

import (
	"encoding/json"
	"net/http"
)

type Engine struct {
	Cc    int `json:"cc"`
	State string
}

func (e *Engine) start() {
	e.State = "started"
}

type Car struct {
	Engine
}

func main() {

	car := Car{
		Engine{
			Cc: 1500,
		},
	}

	car.Engine.start()

	http.HandleFunc("/car", func(res http.ResponseWriter, req *http.Request) {
		carJson, err := json.Marshal(car)
		if err != nil {
			panic(err)
		}
		res.Write(carJson)
	})

	http.ListenAndServe(":8000", nil)

}
