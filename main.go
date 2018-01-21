package main

import "fmt"

func main() {
	addFunc := func(terms ...int) (numTerms int, sum int) {
		sum = 0
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)
		return
	}
	fmt.Println(addFunc(2, 3, 4))
}
