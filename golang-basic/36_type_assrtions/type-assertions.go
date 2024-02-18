package main

import "fmt"

// any/interface = kosong
func random() any {
	return "ok"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)
}
