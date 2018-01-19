package main

import "fmt"

func main() {
	fmt.Println("Input a number:")
	var myNumber int
	fmt.Scanln(&myNumber)
	divByFive := (myNumber%5 == 0)
	divByThree := (myNumber%3 == 0)
	switch {
	case divByFive && divByThree:
		fmt.Println("FizzBuzz")
	case divByThree:
		fmt.Println("Fizz")
	case divByFive:
		fmt.Println("Buzz")
	default:
		fmt.Println(myNumber)
	}
}
