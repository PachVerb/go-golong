package main

import (
	"./conv"
	"fmt"
)

func main() {
	c, b := false, true
	i, j := 1, 0
	//布尔类型转数字
	fmt.Println(conv.Btoi(c), conv.Btoi(b)) // 0 ， 1
	//数字转布尔类型
	fmt.Println(conv.Itob(i)) //true
	fmt.Println(conv.Itob(j)) //false
}
