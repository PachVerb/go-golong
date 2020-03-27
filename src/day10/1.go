/*
	计算结果溢出
	算数运算结果，不管是有符号还是无符号类型，超出类型宽度的。都要抛弃高位，直到符合范围。
	如果是有符号类型，对于最左边bit为是 1 的最终结果可能是负数。
*/
package main

import "fmt"

var u uint8 = 255
var i int8 = 127

func main() {
	fmt.Println(u, u+1, u*u) // 255, 256（超出uint8数值范围，最终结果为 0）u*u通常溢出，最终结果1。。
	fmt.Println(i, i+1, i*i) //

	a := [...]int8{1, 2, 3, 4, 40, 50, 100}
	for _, val := range a {
		fmt.Println(val + 100) //从索引3开始·，所有数溢出，最终结果都为负数。这是对于有符号类型。最左边高位是 1的情况。
	}
}
