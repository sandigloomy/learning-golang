package main

import "fmt"

func sayHello(name string) string {
	fmt.Println("Hello", name)
	// Mengembalikan Data
	return name
}

func isMarried(married bool) bool {
	if married {
		fmt.Println("Sudah Menikah")
	} else {
		fmt.Println("Belum Menikah")
	}
	return married
}

func main() {
	sayHello("Sandi")
	sayHello("Salsa")

	status := isMarried(false)
	fmt.Println(status)

}
