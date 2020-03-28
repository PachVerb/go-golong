/*
	布尔类型
*/
package main

import "fmt"

//布尔类型不会隐式转换为 1 0
func main() {
	i := 1
	fmt.Printf("%v\n", i+int(false))
}
