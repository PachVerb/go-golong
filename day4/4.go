/*
    变量

*/
package main

import "fmt"

func main() {
    //var a string, b int, c bool          声明一组为初始化的不同类型的变量，运行出错。
    var a, b, c string  //声明一组相同形同类型变量
    var str, num = "PackVerb", 13
    fmt.Println(a, b, c)     //excute success
    fmt.Println(str, num)   //excute success
}