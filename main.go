package main

import "fmt"

func main() {
	/* Collections in Go */
	mySlice := []int{42, 43, 44, 45}
	mySlice = append(mySlice, 46)
	fmt.Println(mySlice)
}
