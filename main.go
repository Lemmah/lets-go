package main

import "fmt"

func main() {
	mp := messagePrinter{"This is the message"}
	mp.printMessage()
}

type messagePrinter struct {
	message string
}

// A method of the message printer
func (mp *messagePrinter) printMessage() {
	fmt.Println(mp.message)
}
