package main

import "fmt"

func main() {

	type NoKTP string

	var ktpSandi NoKTP = "11111111111"

	var contoh string = "2222222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpSandi)
	fmt.Println(contohKtp)

}
