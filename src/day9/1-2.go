/*
	隐式作用域，词法域对应的作用域范围不太明确。
	流程控制语句，if， for头部的声明语句。也对应一个独立的作用域。称为隐时作用域。
*/
package main

import "fmt"

func main() {

	x := "hello world"

	for _, x := range x { //在for头部的声明的x 变量, 独立与main函数，for语法块显示作用域。形成单独的隐式作用域。
		x := x + 'A' - 'a'
		fmt.Printf("%c\n", x)
	}

	//如下在包级别声明变量y,可以在之前引用。
	fmt.Println(y)
}

var y = "left"
