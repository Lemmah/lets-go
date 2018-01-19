package main

import "fmt"

/* This is the replacement of the while loop in many languages */
func main() {
	i := 0
	for {
		i++
		fmt.Println(i)
		if i > 10 {
			break
		}
	}
}
