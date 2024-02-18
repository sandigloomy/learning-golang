package main

import "fmt"

func main() {
	// Nilainya Tidak Bisa di Ubah setelah di deklarasikan
	const (
		firstname = "sandi"
		lastname  = "salsa"
		value     = 100
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(value)

}
