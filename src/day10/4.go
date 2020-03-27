/*
	fmt包格式化输出Printf使用技巧
*/

package main

import "fmt"

//以输出不同进制数为列
func main() {
	/*
		通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，
		但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。
	*/
	o := 0666                          //八进制数
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	//通过转意字符，获取unicode码点对应的字符面值.
	//字符面值
	ascii := 'a'
	unicode := '国'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q\n", ascii) //97 对应unicode的码点 字符面值 'a'
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)

}
