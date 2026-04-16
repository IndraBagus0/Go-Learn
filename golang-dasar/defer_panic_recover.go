package main

import "fmt"

// defer
func logging() {
	fmt.Println("Logging")
}

// panic and recover
func endApp() {
	fmt.Println("end app")
	msg := recover()
	fmt.Println("Terjadi error", msg)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("error")
	}
}

func main() {

	runApp(true)

}
