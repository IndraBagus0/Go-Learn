package main

import (
	"belajar-golang/database"
	_ "belajar-golang/database"
	"fmt"
)

func main() {
	fmt.Println(database.GetDB())
}
