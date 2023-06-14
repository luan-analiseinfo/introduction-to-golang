package main

import (
	"fmt"
	"time"
)

func event(text string) {
	time.Sleep(time.Second * 1)
	fmt.Println(text)
}

func heavy() {
	for {
		go event("Heavy")
	}
}
func superHeavy() {
	for {
		go event("super heavy")
	}
}

func main() {
	go heavy()
	go superHeavy()
	fmt.Println("Fin")
	select {}
}
