/*
    测试：变量逃逸
    逃逸意味着程序需要额外为变量分配内存空间
*/

package main

import (
    "fmt"
)
var global *string
/*
    如下局部变量a，通过指针赋值给包级别变量global。函数运行完a 是没有被GC回收的。
    同时变量声明时，需要为变量分配堆内存。以便变量a 能够在全局变量global中被访问到.
*/
func g() string {
    a := "hello world"
    global = &a
    return a
}

//非逃逸, f函数在执行完后，变量x的引用为0，GC回收内存空间.
func f() int {
    x := new(int)
    *x = 1
    return *x
}
func main() {
   fmt.Println(g(), global, *global)
   fmt.Println(f())
}