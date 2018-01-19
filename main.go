package main

func main() {
	/* Pass by refrence vs Pass by Value */
	/* Pass by refrence is what we're used to, passing a copy of say a variable into a function, the real variable is not manipulated by the func but rather a copy. */
	/* Pass by value, you pass the memory address of the variable so it can be manipulated from within the function.*/

	/* Examples: */
	var message = "Lemmah"
	passByRefrence(message)
	println(message)
	passByValue(&message)
	println(message)
}

func passByRefrence(message string) {
	message = "James Lemayian" // Does not change anything
}

func passByValue(message *string) {
	*message = "James Lemayian" // The message is changed now!
}
