/*
	基础数据类型。
	整形：
	分有符号和无符号两种类型，两种类型分别对应 int8, int 16, int 32, int 64
	和 无符号 uint8, uint 16, uint32, uint64大小不同类型。

	int16:  -32768 到 +32767

	int == int32
	表示数值范围： -2,147,483,648 到 +2,147,483,647，有符号整数

	int64: -9,223,372,036,854,775,808 到 +9,223,372,036,854,775,807

	uint则是不带符号的，表示范围是：2^32即0到4294967295
*/

package main

import "fmt"

func main() {
	//var a int16 = 32768	//常数溢出
	var b = -32768
	//var c uint16 = -32768	//常数溢出

	var d uint16 = 32767
	var e uint16 = 32768

	fmt.Printf("int32: %d\n", b)
	fmt.Printf("uint16: %d\n", d)
	fmt.Printf("uint16: %d\n", e)
}
