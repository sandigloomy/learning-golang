package main

import "fmt"

func ups() interface{} {
	return "ups"
}

func main() {
	kosong := ups()
	fmt.Println(kosong)
}
