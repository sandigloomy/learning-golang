package main

import (
	"fmt"
	"learn-goalng/41_package/helper"
	// "golang-basic/41_package/helper"
)

func main() {
	result := helper.SayHello("sandi")
	fmt.Println(result)
	fmt.Println(helper.Application)
}
