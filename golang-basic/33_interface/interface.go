package main

import "fmt"

// Interface adalah tipe data Abstract, dia tidak memiliki imlplementasi langsung
// sebuah interface berisikan definisi-definisi method
// biasanya interface digunakan sebagai kontrak

type HasName interface {
	GetName() string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "sandi"}
	sayHello(person)

	animal := Animal{Name: "kucing"}
	sayHello(animal)
}
