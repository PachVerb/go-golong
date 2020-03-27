/*
	uint 8bit位演示
	%b打印二进制格式数字，08表示至少打印八个宽度位，不足用0补充。
*/
package main

import "fmt"

var x uint8 = 1<<1 | 1<<5
var y uint8 = 1<<1 | 1<<2

func main() {
	fmt.Printf("%08b\n", x) // 00100010
	fmt.Printf("%08b\n", y) // 00000110

	//位移运算
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x&^y)
	fmt.Printf("%08b\n", x<<y)
	fmt.Printf("%08b\n", x>>y)
	//为运算符, 对于有符号类型.
	fmt.Println("---------------------")

	var a = -2

	fmt.Printf("%08b\n", a)    //-0000010	//符号为占一位,原本是 0
	fmt.Printf("%08b\n", a>>1) //-0000001	//符号为占一位

	//go语言提供无符号整形.但是程序中我们更倾向与int类型。
	var f = [...]string{"bananer", "apple", "orange"}
	//续输出数组类型。len函数，返回有符号int类型。如果 是 len返回无符号类型，
	//那么for循环条件永远为真，当 i = 0时候， 在减 1，这时候无符号类型，返回最大值。超出slice边界。

	for i := len(f) - 1; i >= 0; i-- {
		fmt.Println(f[i])
	}
	var d uint8 = 1

	fmt.Println(d - 2) //常数溢出, uint类型返回最大值
}
