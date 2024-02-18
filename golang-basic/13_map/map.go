package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "sandi",
		"address": "bandung",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

}
