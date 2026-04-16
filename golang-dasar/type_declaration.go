package main

import "fmt"

func main() {

	type noKtp string

	var ktp noKtp = "1111111"
	var contoh string = "2222222"

	var contohKtp noKtp = noKtp(contoh)

	fmt.Println(ktp)
	fmt.Println(contohKtp)
}
