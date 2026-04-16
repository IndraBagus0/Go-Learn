package main

import "fmt"

func main() {
	var a = 10
	var b = 20

	var pertambahan = a + b
	var pengurangan = a - b
	var perkalian = a * b
	var pembagian = a / b
	var sisaPembagian = a % b

	fmt.Println("a ditambah b sama dengan", pertambahan)
	fmt.Println("a dikurangi b samadengan", pengurangan)
	fmt.Println("a dikali b samadengan", perkalian)
	fmt.Println("a dibagi b samadnegan", pembagian)
	fmt.Println("Sisa pembagin a dan b adalah", sisaPembagian)

	// 	go run .\operator_matematika.go
	// a ditambah b sama dengan 30
	// a dikurangi b samadengan -10
	// a dikali b samadengan 200
	// a dibagi b samadnegan 0
	// Sisa pembagin a dan b adalah 10

	// Augmented Assignments
	var i = 10
	i += 10
	fmt.Println(i)

	// // 	go run .\operator_matematika.go
	// 	20

}
