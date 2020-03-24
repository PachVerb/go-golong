/*
    元祖赋值
    同时初始化一组变量，或者交换两个变量值

*/
package main

import (
    "fmt"
)

func main() {
    var a = [...]int{1, 2, 3, 4, 5}
    fmt.Println(a)
    a[1], a[2] = a[2], a[1]    //交换 数组 元素 1，2位置 值
    fmt.Println(a)  //[1, 3, 2, 4, 5]
}