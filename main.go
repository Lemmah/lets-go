package main

import "fmt"

func main() {
	// It's recommended to create objects on the heap
	// On the heap, we refer to a refrence type
	foo := new(myStruct)
	foo.myField = "bar"
	println(foo.myField)
	// Confirm that it's a refrence type
	fmt.Println(foo)
	bar := &myStruct{"foo"}
	fmt.Println(bar)
}

/* Struct is a custom type */
type myStruct struct {
	myField string
}
