package main

func main() {
	/* Variatic params must always come last in the list of args */
	sayHelloVariatic("Hello", "From", "Lemmah", "James")
}

func sayHelloVariatic(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}
