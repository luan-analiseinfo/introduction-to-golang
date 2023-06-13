package main

import "fmt"

func AnyThing(any interface{}) {
	fmt.Println(any)
}

func main() {
	fmt.Println("vim-go")
	AnyThing(2.44)
	AnyThing("Luan")
	AnyThing(struct{}{})

	mymap := make(map[string]interface{})

	mymap["name"] = "LASFJLSFJLKDSJF"
	mymap["age"] = 10

	fmt.Println(mymap)

}
