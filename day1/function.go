//函数
//go语言中，函数可以没有参数或者接受多个参数
// 传入参数类型，在变量名之后
//主函数不传入参数
package main

import "fmt"

//声明带参数的函数
func add(x int, y int) int {
    return x + y
}

//求积
func product(x, y int) int {    //函数括号后面的int， 表示申明返回int 类型的product函数
    return x * y;
}

//多返回值
func swap(x, y string, r int) (string, string, int) {  //函数括号后面需要为函数指定返回值类型，对应参数个数(这里是对传入参数的原样返回比)
    return x, y, r
}

//直接返回
func split(sum int) (x, y int) {
     x = sum * 4 / 2
     y = sum - 1

     return     //不带参数返回
}

func main(){
    //函数调用
    fmt.Println(add(42, 13))        // --> 55
    fmt.Println(product(3, 4))  // --> 12

    //使用 swap
    a, b, r := swap("hello", "world", 666);
    fmt.Println(a, b, r);      // --> a:hello b:world， r: 666

    //split
    fmt.Println(split(3))       // 6 2
    //解构返回 , 尝试
    //     a, b := split(3)
    //
    //     fmt.Println(a, b);  //失败
}

