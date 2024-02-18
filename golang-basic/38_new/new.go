package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Enggland"
	fmt.Println(alamat1) // Akan Berubah
	fmt.Println(alamat2)
}
