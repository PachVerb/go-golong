/*
	float类型极限
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	//float32能够表示整数的bit位为23. 其余位置用于表示符号和指数。超出范围出现误差。

	var f float32 = 16777216 //该数的二进制位数为25.

	fmt.Println((f + 1) == f) //true ，超出float32能够表示整数的bit位数。出现误差

	//格式化输出浮点数

	for x := 0; x < 20; x++ {
		fmt.Printf("x=%d e^x = %8.3f\n", x, math.Exp(float64(x))) //这里的 8.3表示含义，输出字符宽度为 8 ,浮点数精度为3
	}

}
