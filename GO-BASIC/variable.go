package main

import "fmt"

func main() {

	var firstName string = "John"
	lastName := "Dhoe"
	lastName = "Abdul"
	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// * Multiple Variable

	var first, second, third string = "satu", "dua", "tiga"
	buah1, buah2, buah3 := "Apel", "Mangga", "Duren"

	fmt.Printf("%s %s %s!\n", buah1, buah2, buah3)
	fmt.Printf("%s %s %s!\n", first, second, third)

	// * Variable underscore

	_ = "Variable keranjang sampah"

}
