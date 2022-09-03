package main

import "fmt"

func main() {
	// * variable Konstanta di go sama kaya di JS artinya nilai di variable tersebut tidak dapat diubah, hanya dapat diinisialisasi sekali di awal pembuatan

	const name string = "John"
	// error -> name := "Joko"

	fmt.Print("Halo ", name, "!\n")

}