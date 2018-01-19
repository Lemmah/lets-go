package main

import "fmt"

func main() {
	/* Collections in Go */
	mySlice := make([]int, 100)
	mySlice[0] = 43
	mySlice[1] = 44
	mySlice[2] = 45
	mySlice = append(mySlice, 46)
	fmt.Println(mySlice)
}
