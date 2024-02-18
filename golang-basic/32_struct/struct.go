package main

import "fmt"

/*
- Struct adalah sebuah template data yang  digunakan untuk menggabungkan  nol atau lebih tipe data lainnya dalam satu kesatuan
- struct biasanya repsentasi data dalam aplikasi yang kita buat
- data struct disimpan dalam field
- sederhananya struct adalah kumpulan dalam field
*/

type Customer struct {
	Name, Address string
	Age           int
}

// Struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var sandi Customer

	sandi.Name = "sandi"
	sandi.Address = "Bandung"
	sandi.Age = 18

	fmt.Println(sandi)

	// struck Literal
	salsa := Customer{
		Name:    "salsa",
		Address: "Bandung",
		Age:     17,
	}
	fmt.Println(salsa)

	sansal := Customer{"sansal", "bandung", 178}
	fmt.Println(sansal)

	praya := Customer{Name: "praya"}
	praya.sayHello("sandi")
	salsa.sayHello("sandi")

}
