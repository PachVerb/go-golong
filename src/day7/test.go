/*
	导入convert目录下的两个同名包 conv, tempconv
	进行华氏温度和摄氏温度转换.
*/
package main

import (
	"./convert"
	"./unit_conv"
	"fmt"
)

//使用tempconv.go中变量
func main() {
	//获取摄氏温度 absolutezeroC: -273.15
	fmt.Println(tempconv.AbsoluteZeroC)
	//将摄氏温度转换为华氏温度
	//摄制温度
	fmt.Printf("F: %s\n", tempconv.CToF(tempconv.AbsoluteZeroC)) //转换后 -459.699999

	fmt.Println(conv.A.String())  //2kg
	fmt.Println(conv.Ga.String()) //100G
}
