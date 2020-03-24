/*
   在函数内部 可以 := 声明变量。这种声明变量使用场景是，不需要显示指定变量类型，根据表达式推导出变量类型。因其简短而广泛使用。
   var声明变量通常是需要显示指定类型。
   := 是变量声明语句，区别与赋值 = 语句.
*/
package main

import (
    "fmt"
)
var sky, adress, state = 25, "北京", '晴'      //包一级变量
func main() {
    //简短变量声明初始化一组变量
    //函数以声明和内部同名变量，内部简短声明语句会从新声明。
    sky, state := 23.5, "良"
    //简短变量声明语句，所声明变量并不都是刚声明的。对于在当前词法作用域中以声明变量，接下来只是赋值操作
    sky, adress := 24, "大连"    //sky处赋值
    //至少存在一个新变量，否则编译错误
    //sky, adress := 30, "四川"     //no new variables on left side of :=
    fmt.Println(sky, state)
    fmt.Println(sky, adress, state)
}