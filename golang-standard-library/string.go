package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Indra Bagus", "Indra"))
	fmt.Println(strings.Split("Indra Bagus", " "))
	fmt.Println(strings.ToLower("Indra Bagus"))
	fmt.Println(strings.ToUpper("Indra Bagus"))
	fmt.Println(strings.Trim("      Indra Bagus      ", " "))
	fmt.Println(strings.ReplaceAll("Indra Bagus Indra Syah", "Indra", "Putra"))
}
