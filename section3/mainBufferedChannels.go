package main

import "fmt"

func main() {
	c := make(chan int)

	var numbers []int

	for i := 0; i < 200; i++ {
		numbers = append(numbers, i)
	}

	go func() {
		for _, name := range numbers {
			c <- name
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

}
