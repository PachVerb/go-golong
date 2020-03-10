//命令行参数
//package main
//
// import (
//     "fmt"
//     "os"    //包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。
// )
//
// func main() {
//     var s, sep string
//     for i := 1; i < len(os.Args); i++ {
//         s += sep + os.Args[i]
//         sep = ''
//     }
//     fmt.Println(s)
// }

package main

import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}