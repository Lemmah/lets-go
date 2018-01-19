package main

import "fmt"

func main() {
	var myNumber int
	fmt.Println("Input a number: ")
	fmt.Scanln(&myNumber)
	divByFive := (myNumber%5 == 0)
	divByThree := (myNumber%3 == 0)
	if divByFive && divByThree {
		fmt.Println("FizzBuzz")
	} else if divByFive {
		fmt.Println("Buzz")
	} else if divByThree {
		fmt.Println("Fizz")
	} else {
		fmt.Println(myNumber)
	}
}
