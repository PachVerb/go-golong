/*
	数字转字符串
*/

package main

import (
	"fmt"
	"strconv"
)

var num = 123
var num2 = -123

func main() {
	a := fmt.Sprintf("%d", num) //type-> string 123
	b := strconv.Itoa(num)      //type-> string 123

	fmt.Println(a, b)

	//FormatInt和FormatUint函数可以用不同的进制来格式化数字：
	fmt.Println(strconv.FormatInt(int64(num), 8))
	fmt.Println(strconv.FormatUint(uint64(num2), 2))
	//相比之下，fmt.格式化输出中 参数 %b, %x...等参数比Fortmat更方便

	//字符串转整数
	fmt.Println(strconv.Atoi("123"))              //123
	fmt.Println(strconv.ParseInt("-123", 16, 32)) //291
	fmt.Println(strconv.ParseUint("123", 16, 32))

	//fmt.Scanf来解析输入的字符串和数字

}
