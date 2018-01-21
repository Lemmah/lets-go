package main

func main() {
	// Creating an object in the local memory stac
	foo := myStruct{}
	foo.myField = "bar"
	println(foo.myField)
}

/* Struct is a custom type */
type myStruct struct {
	myField string
}
