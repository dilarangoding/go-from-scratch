package main

import "fmt"

func main(){

	// * Aritmatika
	var angka = (((2+9) % 2) * 5 -1 ) / 2

	// * Perbandingan
	isEqual := (angka == 2)

	fmt.Printf("Nilai %d (%t) \n",angka, isEqual)
	// * Operator Logika
	left := false
	right := true

	leftAndRight := left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	leftOrRight := left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	leftReverse := !left
	fmt.Printf("!left \t(%t) \n", leftReverse)
}