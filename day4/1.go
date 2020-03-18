package main    //.go源文件以包声明语句开始。

import (            //import语句导入依赖包
    "fmt"
)
/*
   包一级变量，常量，函数声明处。
   注: 函数和函数内部变量必须先声明后使用。
*/

//不带操作对象
//无·tag switch == switch true
func Signum(x int) int {
    switch {
    case x > 0:
        return +1
    default:
        return 0
    case x < 0:
        return -1
    }
}

func main(){
    fmt.Println(Signum(2))
}