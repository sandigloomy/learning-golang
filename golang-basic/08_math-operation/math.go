package main

import "fmt"

func main() {

	fmt.Println("Math Operation")

	var a = 30
	var b = 20
	var d = 5
	var c = a + b - d
	fmt.Println(c)

	// Augmente Assignments
	var i = 10
	i *= 10 // i = i + 10
	fmt.Println(i)

	i += 10 // i * i + i
	fmt.Println(i)

	// Unary Operator
	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j++
	fmt.Println(j)

	j-- // j = j - j
	fmt.Println(j)
}
