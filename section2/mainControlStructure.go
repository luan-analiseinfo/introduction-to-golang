package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	f := true
	flag := &f

	if flag != nil {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	arry := []string{"luan", "lucas", "github"}

	for i, s := range arry {
		fmt.Println(i, " ", s)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "Luan"
	mymap["age"] = 18

	for k, v := range mymap {
		fmt.Printf("Key: %s and Value %v\n", k, v)
	}

}
