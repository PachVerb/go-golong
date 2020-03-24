/*
    通过命令行参数为变量设置值
*/
package main

import (
    "fmt"
    "flag"
    "strings"
)

/*
    程序描述:
    调用flag.bool会创建一个新bool类型命令行参数标志变量.
    Bool函数第一个参数 是 命令行参数名字
    第二个参数是值
    第三个是参数说明

    类似 调用 flag.string创建bool类型命令行参数变量, 同样包含命令行参数名， 变量值, 变量描述
    
    变量n, sep对应变量指针.
*/
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func  main() {
   flag.Parse()    //更新设置的命令行参数
   fmt.Print(strings.Join(flag.Args(), *sep))
   if !*n {
       fmt.Println()
   }
}