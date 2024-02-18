package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "sandi"
	names[1] = "Maulana"
	names[2] = "Fauzan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(values)

	var nama = [...]string{
		"sandi",
		"maulana",
		"fauzan",
	}

	fmt.Println(nama)
	fmt.Println(len(nama))
	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	// Function Array
	// len(array)
	// array[index]
	// array[index] = values

}
