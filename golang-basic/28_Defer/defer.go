package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runApplication() {
	// Setelah Selesai Memanggil Function runApplication
	// defer akan mengeksekusi setelah function runApplication
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
