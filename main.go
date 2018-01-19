package main

import "fmt"

func main() {
	/* Collections in Go */
	myArray := [...]int{42, 43, 44, 45}
	mySlice := myArray[:]
	mySlice = append(mySlice, 46)
	fmt.Println(myArray)
	fmt.Println(mySlice)
}
