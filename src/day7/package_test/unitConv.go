package main

import (
	"fmt"
	//"os"
	"strconv" //该包用于基本数据类型之间的字符串转换。
)

func main() {
	f, err := strconv.Atoi("-23") //Atoi，字符串数字转十进制数字
	fmt.Println(f, err)           // 23

	//ParseBool，ParseFloat，ParseInt和ParseUint字符串转基本数据类型
	a, err := strconv.ParseBool("true")        //true
	b, err := strconv.ParseFloat("3.1415", 64) //该方法，第二个参数用于指定解析宽度。float64
	c, err := strconv.ParseInt("-2", 5, 32)    //该方法，第二位设置基数，第三位是解析精度。
	d, err := strconv.ParseUint("2", 5, 32)    //该方法使用同parseInt一样, 不同的是 parseuint转换无符号字符串数字。

	fmt.Println(a, b, c, d)
}
