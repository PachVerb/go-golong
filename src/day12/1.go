/*
	字符串: len, s[i]测试
*/

package main

import (
	"fmt"
	"os"
)

var s = "hello, world"

func main() {
	fmt.Fprintf(os.Stdout, "%d\n", len(s))
	fmt.Fprintf(os.Stdout, "s[0]:%d s[7]:%d\n", s[0], s[7]) //返回 0，7位置的字节值。104， 119。这是对应unicode吗点的字符。
	//fmt.Fprintf(os.Stdout, "%d", s[14])	//这里访问超出字符索引范围的序列导致panic错误.

	//非ascll字符要多个字节
	var st = "我是, PachVerb"
	fmt.Println(len(st)) //字节数 16. 寿命文本字符串处，占非一个字符
	fmt.Println(st[0:6]) //s[i:j]表示字符串切片. 该切片参数中, 边界0=<i, j < len[s]. 并且不包括边界j. 从前面结果中，我们可以得知。文本字符"我是"占用六个字节.

	// = .
	var t string
	t = s
	fmt.Println(t) //hello world

	// +
	fmt.Println(s[0]) //104
	s += ",I am new PachVerb!"
	fmt.Println(s)    //s 原来字符值被改变.
	fmt.Println(s[0]) //104

	//s[i]， 测试改变字符序列处的值.
	//s[0] = "L"	//Error: can not assign to s[0], 该字节序列处不能被分配.
}
