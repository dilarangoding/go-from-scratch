package main

import "fmt"

func main() {

	// * Tipe Data Numerik Non-Desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -2332233

	fmt.Printf("Bilangan positif : %d\n", positiveNumber)
	fmt.Printf("Bilangan negatif : %d\n", negativeNumber)

	// * Tipe Data Numerik Desimal

	decimalNumber := 3.15

	fmt.Printf("Bilangan Desimal : %f\n", decimalNumber)
	fmt.Printf("Bilangan Desimal : %.2f\n", decimalNumber)

	// * Tipe Data Boolean

	var exist bool = true
	exist2 := false
	fmt.Printf("Exist? %t \n", exist)
	fmt.Printf("Exist? %t \n", exist2)

	// * Tipe Data String

	message := "Halo"
	var message2 = `Gue "John Dhoe".
		lagi belajar bahasa Go 
		biar edgy aje
	`
	fmt.Printf("message : %s \n", message)
	fmt.Println(message2)

	// * Nila nil & Zero Value
	// * Semua tipe data memiliki zero value(nilai default tipe data)
	// * Nil dengan zero value berbeda -> nilai nil benar benar kosong
	
}