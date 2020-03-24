/*
    new 函数声明变量

*/

package main

import (
    "fmt"
)
//new函数声明变量和普通变量声明方式无异
func newInt() *int {
     return new(int)
}
//普通声明
func NewInt() *int {
    var dumy int
    return &dumy
}

//每次调用new 返回新地址, 同一类型多次new, 是不同的变量。

func New() bool {
     p := new(int)
     q := new(int)

     return p == q
}

//基于上面情况，也有列外,类型长度为0. 有可能相同地址
func typeNew() bool {
    p := new(struct{})
    q := new([0]int)

    return *p == *q
}
func main() {
    p := new(int)
    fmt.Println(*p)  //0
    *p = 2
    fmt.Println(*p)    //2

    fmt.Println(newInt())
    fmt.Println(NewInt())

    fmt.Println(New())  //false

    fmt.Println(typeNew())
}