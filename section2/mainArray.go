package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	names := []string{"hi", "my", "name"}

	arr2 := append(names, "is", "Luan", "!")
	fmt.Println(arr2, numbers)
}

/*func main() {
	a := 1
	b := 2

	fmt.Println(a + b)
}
*/
