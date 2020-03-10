package main

import (
	"fmt"
)

func str() string {
	var a, b = "hello", "world"

	return a +
		b //再加号前加分号导致出错 ,在加号后不会出错.
}
func main() { //函数体后面大括号紧随，不能独占一行，否则影响编译
	fmt.Println("hello, 世界")

	fmt.Println("return a + b:", str)
}
