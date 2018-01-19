package main

import "fmt"

func main() {
	fmt.Println("Input a number:")
	var myNumber int
	fmt.Scanln(&myNumber)
	fizzOrBuzz := divisibilityTest(myNumber)
	switch fizzOrBuzz {
	case 1:
		fmt.Println("FizzBuzz")
	case 2:
		fmt.Println("Buzz")
	case 3:
		fmt.Println("Fizz")
	default:
		fmt.Println(myNumber)
	}
}

func divisibilityTest(number int) int {
	if number%5 == 0 && number%3 == 0 {
		return 1
	} else if number%5 == 0 {
		return 2
	} else if number%3 == 0 {
		return 3
	}
	return 0
}
