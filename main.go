package main

import "fmt"

func main() {
	sayHello("Lemmah")
}

func sayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}
