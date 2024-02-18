package min

import "fmt"

func SayHello(name string) string {
	return "Hello" + name
}

func main() {
	fmt.Println(SayHello("sandi"))
}
