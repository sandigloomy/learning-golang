package main

import (
	"fmt"
)

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke", counter)
	}
	fmt.Println("Selesai")

	names := []string{"sandi", "maulana", "fauzan"}

	for i := 0; i < len(names); i++ {
		fmt.Println("perulangan ke", i)
	}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	// perulangan dengan break

	for i := 1; i <= 10; i++ {

		if i > 5 {
			break
		}

		fmt.Println("perulangan ke", i)
	}

	// Perulangan dengan continue

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke", i)
	}

}
