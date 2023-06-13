package main

import "fmt"

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("driving...")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	c := Car{
		Age:     18,
		ModelNo: 1999,
		Name:    "chevy",
	}

	c.Print()
	c.Drive()
	fmt.Println(c.GetName())

}
