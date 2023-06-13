package main

import (
	"fmt"
)

type CarInterface interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

func (l Lambo) Drive() {
	fmt.Println("Lambo on the move ")
	fmt.Println(l.LamboModel)
}

func (l Lambo) Stop() {
	fmt.Println("Lambo stopped ")
	fmt.Println(l.LamboModel)
}

type Chevy struct {
	ChevyModel string
}

func (l Chevy) Drive() {
	fmt.Println("Chevy on the move ")
	fmt.Println(l.ChevyModel)
}

func (l Chevy) Stop() {
	fmt.Println("ChevyModel stopped ")
	fmt.Println(l.ChevyModel)
}

func drive(carInterface CarInterface) {
	carInterface.Drive()
}

func main() {
	cars := []CarInterface{Lambo{"Gallard"}, Chevy{"C369"}}

	for _, car := range cars {
		drive(car)
	}
}
