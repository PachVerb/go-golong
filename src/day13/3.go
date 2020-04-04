/*
	获取每一个unicode字符
	unicode/utf-8包使用
*/

package main

import (
	"./coslib"
	"fmt"
	"unicode/utf8"
)

var str = "hello, 世界"

func main() {
	//通常使用str[index]是无法获取字符的，获取的而是utf-8编码值
	fmt.Println(len(str))                    //13
	fmt.Println(utf8.RuneCountInString(str)) //9 	这是返回字符长度

	//使用uft-8解码器,获取每个unidcode值
	//for len(str) > 0 {
	//	r, size := utf8.DecodeRuneInString(str)
	//	fmt.Printf("%c\t%v\n",r, size)
	//	str = str[size:] //每循环一次，更新str值。如果不添加这一步会死循环。
	//}
	//也可以使用range。range底层会对utf-8进行隐式解码.
	for i, r := range str {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	//使用自定义封装包统计字符长度
	//method 1
	fmt.Println(count.Count(str)) // 9
	//method 2
	fmt.Println(count.Counts(str)) //9

	//错误的utf-8编码对应 unicode转义
	var Eres = "\uFFFD"

	fmt.Println(Eres)

	//string接受到[]rune的类型转换，将一个UTF8编码的字符串解码为Unicode字符序列：
	s := "プログラム"
	fmt.Printf("% x\n", s) //e3 83 97 e3.....
	r := []rune(s)
	fmt.Printf("%x\n", r) //[30d7, 30ed, 30bo, 3039, 3030] 该序列是unicode字符序列
	//对p[]rune类型unicode序列进行切片和转换为string,会进行utf-8编码
	fmt.Println(string(r))

	//整数转换为字符串，则转换包含unicode编码的utf-8字符串。
	fmt.Println(string(65)) // "A", not "65"
	fmt.Println(string(97))
	fmt.Println(string(0x4eac)) // "京"

}
