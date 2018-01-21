package main

func main() {
	// It's recommended to create objects on the heap
	foo := new(myStruct)
	foo.myField = "bar"
	println(foo.myField)
}

/* Struct is a custom type */
type myStruct struct {
	myField string
}
