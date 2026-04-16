package main

import "fmt"

func main() {
	name := "Indra Bagus"
	fmt.Println(name)

	name = "Indra Bagus Syah"
	fmt.Println(name)

	var (
		firstName  = "Indra"
		middleName = "Bagus"
		lastName   = "Syah"
	)

	fmt.Println(firstName, middleName, lastName)
}
