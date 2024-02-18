package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// Mengkonversi function ke dalam variable
	goodbye := getGoodBye
	ex := getGoodBye

	fmt.Println(goodbye("Sandi"))
	fmt.Println(ex("salsa"))
}
