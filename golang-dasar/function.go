package main

import "fmt"

func sayHello(name string, umur int) string {
	// fmt.Println("Hello", name, "Umur", umur)
	hello := "Hi " + name + " Umur " + fmt.Sprintf("%d", umur)
	return hello
}

func helloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hi", filter(name))

}

func filteredName(name string) string {

	if name == "Anjing" {
		return "..."
	} else {
		return name
	}

}

func main() {
	name := "indra"
	// umur := 22
	// result := sayHello(name, umur)
	// fmt.Println(result)
	helloWithFilter(name, filteredName)
}
