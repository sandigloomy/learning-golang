package main

import "fmt"

func main() {

	fmt.Println(true)
	fmt.Println(false)

	// Operasi Boolean

	// && = Dan
	// || = Atau
	// ! = Kebalikan

	var nilaiAkhir = 90
	var absensi = 81

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80

	var lulus = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

}
