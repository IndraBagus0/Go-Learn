package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {

		// break
		// if i == 5 {
		// 	break
		// }

		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}
