package main

func main() {
	/* variable datatypes in go */
	var myInt int
	myInt = 43
	println(myInt)
	secondInt := 44 // the shorthand version
	println(secondInt)
	var myFloat float32 = 32.
	println(myFloat)
	myFloatTwo := 32.
	println(myFloatTwo)
	myString := "Hello Go!"
	println(myString)
	myComplex := complex(3, 2)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))
}
