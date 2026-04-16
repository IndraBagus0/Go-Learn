package main

import "fmt"

func main() {
	name := "Bagus"

	if name == "Indra" {
		fmt.Println("haloo nama saya indra")
	} else if name == "joko" {
		fmt.Println("halo joko")
	} else {
		fmt.Println("halo nama saya", name)
	}

	length := len(name)

	if length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("Nama Suadah Benar")
	}
}
