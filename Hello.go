package main

import "fmt"

// Variable
var world = "world"
var text, text2 = "Luqmanul", "Hakim"
var space string

func main() {
	// Shorthand variable hanya bisa didalam fungsi main
	number := 123+321
	var hasil int
	hasil = 23*2
	
	fmt.Printf("hello, " + world + "\n")
	fmt.Println(text, text2)
	fmt.Println(space)

	fmt.Println(number)
	fmt.Println(hasil)
	
	hasil++
	fmt.Println(hasil)
	hasil--
	fmt.Println(hasil)
}