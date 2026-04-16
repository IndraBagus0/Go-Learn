package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// 	# go run .\conversion.go
	// 32768
	// 32768
	// -32768

	// kenapa nilai16 minus? karena nilai maksimun int16 adalah 32768

	var name = "Indra Bagus"
	var e uint8 = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
	// 	# go run .\conversion.go
	// 	Indra Bagus
	// 73
	// I
}
