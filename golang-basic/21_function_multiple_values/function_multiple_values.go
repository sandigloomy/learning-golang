package main

import "fmt"

func sayHello() (string, string) {
	return "sandi", "salsa"
}

func main() {
	firtsName, lastName := sayHello()

	fmt.Println(firtsName, lastName)
}
