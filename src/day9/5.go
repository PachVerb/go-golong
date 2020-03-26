/*
	算数运算符
	取 %运算，用于整数间运算。其符号有被取模数符号决定。
	/运算
*/
package main

import "fmt"

var a, b, c, d = 5, -5, 3, -3

func main() {
	//%
	fmt.Println(a % c) //2
	fmt.Println(a % d) //2
	fmt.Println(b % c) //-2
	fmt.Println(b % d) //-2

	//  / 运算，余数问题. 如果全为整数，不管是否能处境. 截断余数。 否则保留余数
	a, b, c, d := 1, 2, 4.2, 2.0
	fmt.Println(a / b) //1
	//fmt.Println(b/d)	//不同类型无法运算
	fmt.Println(c / d) //2.04...
}
