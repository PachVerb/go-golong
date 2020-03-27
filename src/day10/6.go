/*
	math包中的特殊值
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var z float64                      //为初始化，赋值 0
	fmt.Println(z, -z, 1/z, -1/z, z/z) //0, -0, inf, -inf, nan

	//math.isNaN, math.NaN
	//math.isNaN测试一个数是否是非数
	fmt.Println(math.IsNaN(z)) //false
	//math.NaN返回非数对应的值
	fmt.Println(math.NaN()) //NaN

}
