package helper

// diawali dengan huruf besar agar bisa di akses diluar
var version = "1.0.0"      //tidak bisa di akses diluar package
var Application = "golang" // bisa di akses

func SayHello(name string) string {
	return "Hello " + name
}
