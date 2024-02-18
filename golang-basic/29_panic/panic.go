package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	// Bila Aplikasi Terjadi Error defer akan tetap mengeksekusi
	defer endApp()
	if error {
		panic("Error")
	}
}

func main() {
	// jika ini true maka error tapi defer tetap jalan
	runApp(false)
}
