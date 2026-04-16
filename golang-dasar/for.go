package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 11 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Perulnagan ke ", counter)
	}

	fmt.Println("Done")

	names := []string{"Indra", "Bagus", "Syah", "Putra"}
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
}
