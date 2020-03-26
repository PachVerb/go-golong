/*
	词法域决定了作用域的大小，
	在函数内词法与嵌套
*/
package main

import "fmt"

func main() {
	x := "hello!"
	//a := 'h'	//变量声明为使用表错
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c\n", x)
		}
	}
	//fmt.Printf("%d", a)
}
