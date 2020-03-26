/*
	位移运算符。
	可以利用左移和右移运算中，每次移动都表示一个数的 2 次幂这个特性，来作为某些乘法和除法运算的小技巧。
*/
package main

import "fmt"

var c, b, e, z = 60, 13, -60, -13

func main() {
	//&运算：二元运算符。两数二进制同时为1，对应位为1， 否则为零
	//c:111100, b: 1101 测试组合， 001100
	fmt.Println(c & b) //12
	//|运算: 两者两个数同时为1或者1个为1一个不为1的时候会保留。
	//c: 111100, b: 001101 |=》, 111101 如果两数，二进制数位长度不同，二进制数前可以用零补位.
	fmt.Println(c | b) //61
	//^按位异域运算符：当双目运算符使用时, 两数对应二进制位不同时，取 1. 否则保留相同值其一。
	//c:111100, b: 001101 ^=>, 110001
	fmt.Println(c ^ b) //49
	//^当一元运算符使用,表示取反. 经测试发现取反规则：原数 + 1 取 + -
	fmt.Println(^c) //-61
	fmt.Println(^b) //-14
	fmt.Println(^e) //59
	fmt.Println(^z) //12
	//<< 左移运算符
	fmt.Printf("left:%08b\n", c) //二进制格式化输出数, 00111100
	fmt.Printf("left-source:%d\n", c<<1)
	fmt.Printf("%08b\n", c<<1) //左移动一位，01111000. 左移，右边空出的位置用零补位.
	fmt.Printf("left-source:%d\n", c<<2)
	fmt.Printf("%08b\n", c<<2) //左移动两位, 11110000
	fmt.Printf("left-source:%d\n", c<<3)
	fmt.Printf("%08b\n", c<<3) //左移动三位, 111100000

	//>>右移运算符.
	fmt.Printf("right:%08b\n", c)
	fmt.Printf("right-source:%d\n", c>>1)
	fmt.Printf("right-1:%08b\n", c>>1) //右移动一位,空出高位补零，00011110
	fmt.Printf("right-source:%d\n", c>>2)
	fmt.Printf("right-2:%08b\n", c>>2) //右移动两位, 00001111
	fmt.Printf("right-source:%d\n", c>>3)
	fmt.Printf("right-3:%08b\n", c>>3) //三位, 00000111
	fmt.Printf("right-source:%d\n", c>>4)
	fmt.Printf("right-4:%08b\n", c>>4)
	fmt.Printf("right-source:%d\n", c>>5)
	fmt.Printf("right-5:%08b\n", c>>5)
	fmt.Printf("right-source:%d\n", c>>6)
	fmt.Printf("right-6:%08b\n", c>>6)
	fmt.Printf("right-source:%d\n", c>>7)
	fmt.Printf("right-7:%08b\n", c>>7)

	fmt.Println(1 << 2)
}
