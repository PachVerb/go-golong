/*
	算数逻辑二元运算符，要求两边操作数类型相同。对于类型不同的，需要显示通过某类型关键字 T(x)进行转换。这个操作的目的告诉编译器，数据类型，
	以及内部结构，该类型转换操作不会改变变量值。

	对于每种类型T，如果转换允许的话，类型转换操作T(x)将x转换为T类型。

*/
package main

import "fmt"

//不同类型的算数运算
func main() {
	var a int8 = 13
	var b int16 = 15

	//fmt.Println(a+b)	//invalid operation: mismatched types int8 and int16, 类型不匹配。

	//显示转换，解决不同类型的算数运算
	fmt.Println(int(a) + int(b)) // 28
	/*
		许多整形数之间的相互转换并不会改变数值；它们只是告诉编译器如何解释这个值。
		但是对于将一个大尺寸的整数类型转为一个小尺寸的整数类型，或者是将一个浮点数转为整数，可能会改变数值或丢失精度：
	*/

	//浮点数, 转换int类型丢失精度
	f := 3.1415
	fmt.Println(int(f)) //丢失精度, 3
	//大尺寸int类型转小尺寸int类型。数值改变
	var z int32 = 16
	var r int64 = 499

	fmt.Println(int8(z), int8(r)) //16, -13

	//还有一种情况是超出目标类型范围的转换，这种情况依赖转换具体实现。
	i := 1e100
	fmt.Println(int(i)) // -92233720....，可以发现转换类型后的值和原来的数值相差很大。

}