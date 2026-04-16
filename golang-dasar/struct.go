package main

import "fmt"

type Customer struct {
	Name, Addres string
	age          int
}

func main() {
	var indra Customer
	indra.Name = "Indra"
	indra.Addres = "Indonesia"
	indra.age = 22

	fmt.Println(indra)
}
