package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Sandi", spamFilter)

	filter := spamFilter
	// sayHelloWithFilter("Anjing", filter)
	sayHelloWithFilter("Anjing", filter)
}
