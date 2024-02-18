package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	sandi := Man{"sandi"}
	sandi.Married()
	fmt.Println(sandi)
}
