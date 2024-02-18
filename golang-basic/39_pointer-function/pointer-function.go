package main

import "fmt"

// Address adalah struktur yang merepresentasikan alamat dengan kota, provinsi, dan negara.
type Address struct {
	City, Province, Country string
}

// ChangeCountryToIndonesia adalah fungsi yang mengubah nilai negara pada alamat menjadi "Indonesia".
// Fungsi ini menerima pointer ke Address sebagai argumen.
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia" // Mengubah nilai Country dari alamat yang ditunjuk oleh pointer
}

func main() {
	// Membuat variabel address dengan tipe data Address
	var address Address = Address{}

	// Memanggil fungsi ChangeCountryToIndonesia dengan memberikan alamat variabel address (&address)
	ChangeCountryToIndonesia(&address)

	// Mencetak nilai address setelah diubah
	fmt.Println(address)
}
