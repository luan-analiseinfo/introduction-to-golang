package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c <-chan int, quits <-chan struct{}) {
	//switch case
	//select for channel async
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("recived")
		case <-quits:
			fmt.Println("Quit")
			os.Exit(0)
		}
	}
}

func main() {
	c := make(chan int)
	quits := make(chan struct{})

	go func() {
		c <- 1
		quits <- struct{}{}
	}()

	go Select(c, quits)
	select {}
}
