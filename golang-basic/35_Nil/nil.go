package main

import "fmt"

func newMap(name string) map[string]string {
	//
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// Jika data kosong
	data := newMap("sandi")

	if data == nil {
		// jalankan ini
		fmt.Println("Data Masih Kosong")
	} else {
		fmt.Println(data["name"])
	}
}
