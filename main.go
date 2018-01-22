package main

import "fmt"

func main() {
	mp := enhancedMessagePrinter{}
	mp.message = "This is the message"
	mp.printMessage()
}

type messagePrinter struct {
	message string
}

// A method of the message printer
func (mp *messagePrinter) printMessage() {
	fmt.Println(mp.message)
}

// Object composition is go's replacement of inheritance
type enhancedMessagePrinter struct {
	messagePrinter
}
