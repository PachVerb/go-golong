/*
	复数
*/
package main

import (
	"fmt"
	"math/cmplx"
)

//构建复数
var a = complex(1, 2) // 1+2i
var b = complex(3, 4) //3+4i
func main() {
	fmt.Println(a) //1+2i
	fmt.Println(b) //3+4i
	//复数乘法运算法，类似多项式乘法。展开过后合并实部和虚部。
	fmt.Println(a * b) //-5+10i
	//加法运算
	fmt.Println(a + b) //4+6i

	//内建real返回实部
	fmt.Println(real(a)) //1
	//内建image返回虚部
	fmt.Println(imag(b)) //4
	//复数比较 ==
	fmt.Println(a == b)

	//求复数平方根
	fmt.Println(cmplx.Sqrt(4)) //2+0i
	fmt.Println(cmplx.Sqrt(-1 + 1i))
	fmt.Println(cmplx.Sqrt(-1)) //0 + 1i

}
