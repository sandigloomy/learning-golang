package main

import "fmt"

func main() {

	nama := "salsa"

	switch nama {
	case "sandi":
		fmt.Println("hallo sandi")
	case "salsa":
		fmt.Println("hallo salsa")
	default:
		fmt.Println("hallo kamu siapa?")
	}

	length := len(nama)

	switch {
	case length > 10:
		fmt.Println("nama sudah benar")
	case length < 10:
		fmt.Println("nama terlalu pendek")
	default:
		fmt.Println("nama tidak ditemukan")
	}

}
