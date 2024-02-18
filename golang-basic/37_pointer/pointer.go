package main

import "fmt"

// Address adalah struktur yang merepresentasikan alamat dengan kota, provinsi, dan negara.
type Address struct {
	City, Province, Country string
}

func main() {
	// Membuat variabel address1 dengan data alamat Bandung, Jawa Barat, Indonesia
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	// Membuat variabel address2 yang merupakan pointer ke address1
	address2 := &address1     // (Copy Value)
	address2.City = "Jakarta" // Mengubah nilai City dari address2
	fmt.Println(address1)     // Ikut Beubah
	fmt.Println(address2)     // Berubah Menjadi Jakarta

	// Mengubah nilai City Province Country dari variable address1
	// address2 = &Address{"Cipadung", "Cibiru", "Indonesia"}
	*address2 = Address{"Bali", "Cibiru", "Indonesia"}

	// Mencetak nilai address1 dan address2
	fmt.Println(address1) // Nilai tidak akan berubah karena address2 adalah salinan nilai dari address1
	fmt.Println(address2)

}
