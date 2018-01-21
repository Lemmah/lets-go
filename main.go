package main

func main() {
	/* Variatic params must always come last in the list of args */
	greetings := []string{"Hello", "From", "Lemmah", "James"}
	sayHelloVariatic(greetings...)
}

func sayHelloVariatic(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}
