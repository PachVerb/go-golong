/*
    指针 var x int
    $x 获取变量指针
    *获取变量的值
    根据变量类型，可以称指针为对应类型的指针.
*/
package main

import "fmt"

func main() {
    var x int = 2
    p := &x //指针
    *p = 7        //修改指针对应的变量的值.
    fmt.Println(x)

    //指针比较
    //只有当指针指向同一变量或者都为nil时，两指针相等.
    var z, y int
    fmt.Println(&z == &y, &z == &z, &y == &y, &z == nil,  &z, &y)

    var t = f()

    fmt.Println(f() == f(), t == t, f(), f())     //false , true

    v := 1
    fmt.Println(incr(&v))     //2
}
//返回局部变量指针
func f() *int {
    v := 1
    return &v
}
//通过指针改变变量值
func incr(p *int) int {
    *p++
    return *p
}