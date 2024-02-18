package main

import "fmt"

func main() {
	fmt.Println("Conversion")

	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var nama = "Sandi"
	var e uint8 = nama[1]
	var eString = string(e)

	fmt.Println(nama)
	fmt.Println(e)
	fmt.Println(eString)
}
