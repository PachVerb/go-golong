/*
	因为Go语言源文件总是用UTF8编码，并且Go语言的文本字符串也以UTF8编码的方式处理，因此我们可以将Unicode码点也写到字符串面值中。

	字符串面值中, 可以插入 以 \开头的转移数据.

*/
package main

import "fmt"

func main() {
	//字符面值插入转义数据
	fmt.Println("hello, world\t i am pachverb")       //\t 制表符
	fmt.Println("今天天气不错! \a 一起来晒太阳吧")                 //\a 响铃
	fmt.Println("春天来啦，\f小树抽出啦绿芽。 导出鸟语花香，呈现出一阵又一阵热闹.") //\f换页

	//字符串值中包含 十六进制，和八进制.
	fmt.Println("i am age\t\123") // 插入转义八进制 \123。
	fmt.Println("i am age\t\377") //注意八进制转移为 \377。 对应十进制 255. 每一个字符对应一个字节。字节类型byte的范围是 255.
	//因此切勿超过 /377
	fmt.Println("i am age\x99") //插入转移十六进制 \x99

	//原生字符串
	s := `hello world\t    
	` //反引号表示纯原生字符串，无转义操作. 包含回车和空格
	fmt.Println(s)

	//原生字符串写入值.
	//这里 通过 + " `这是字符面值`这里字符串"
	s += "`hello\x99` \355"
	fmt.Println(s)
}
