package main

import "fmt"

/* This is the replacement of the while loop in many languages */
func main() {
	i := 0
	for {
		i++
		fmt.Println(i)
		if i > 10 {
			break
		}
	}

	/* Looping through a slice */
	mySlice := []string{"Lemmah", "James", "Go", "Annie"}
	for index, value := range mySlice {
		fmt.Println(index, value)
	}

	/* Looping through a map */
	myMap := make(map[string]string)
	myMap["Go Developer"] = "Lemmah"
	myMap["Heart Beat"] = "Annie"
	myMap["Home Village"] = "Maasai Mara"
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
