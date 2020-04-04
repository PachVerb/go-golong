/**/
package main

import "fmt"

func main() {
	fmt.Println("\u5B57") //字
	fmt.Println("0x2FAC")
	var cstr = "汉字"

	fmt.Println(cstr[1]) //177

	//go字符串，unicode转义输出
	//var cstr = "世界"
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c") //世界
	fmt.Println("\u4e16\u754c")             //世界
	fmt.Println("\U00004e16\U0000754c")     //世界。

	//输出变量实际类型
	var a byte = 'a'
	var crun rune = '国'

	fmt.Printf("%d%T\t\n", a, a)       //ascll 97, type uint8 == byte
	fmt.Printf("%d%T\t\n", crun, crun) //unicode22269 type int32

	//unicode rune类型
	//a rune := '世'
	//转义
	var ca rune = '\u4e16'
	var cb rune = '\U00004e16'

	fmt.Println(ca, cb)
}
