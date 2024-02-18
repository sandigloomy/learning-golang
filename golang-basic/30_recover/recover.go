package main

import "fmt"

func endApp() {
	fmt.Println("End App")

	// jika terjadi panic recover akan memberitahu
	message := recover()
	fmt.Println("Terjadi Error ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("eror ada panic")
	}

}

func main() {
	runApp(true)
}
