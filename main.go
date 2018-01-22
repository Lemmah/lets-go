package main

import "fmt"

func main() {
	foo := newMyStruct()
	foo.myMap["foo"] = "bar"
	fmt.Println(foo)
}

type myStruct struct {
	myMap map[string]string
}

func newMyStruct() *myStruct {
	result := myStruct{}
	result.myMap = map[string]string{}
	return &result
}
