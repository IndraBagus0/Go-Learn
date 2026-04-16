package main

import (
	"fmt"
)

func main() {
	name := "Bagus"

	switch name {
	case "Bagus":
		fmt.Println("halo Bagus")
	case "Indra":
		fmt.Println("halo Indra")
	default:
		fmt.Println("halo", name)

	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Banar")
	}
}
