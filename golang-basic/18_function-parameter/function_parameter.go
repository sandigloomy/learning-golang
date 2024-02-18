package main

import "fmt"

// Funcion With Parameter
func sayHello(firtsName string, lastName string) {
	fmt.Println("Hello", firtsName, lastName)
}

func main() {
	sayHello("sandi", "maulana")
}
