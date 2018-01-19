package main

import "fmt"

func main() {
	/* maps, dictionaries in go */
	myMap := make(map[int]string)
	myMap[1] = "Lemmah"
	myMap[2] = "James"
	fmt.Println(myMap)
}
