//导出名, 使用导出包的方法，首字母大写为导出状态，小写未导出
package main

import (
    "fmt"
)

func main() {
    fmt.Println("ok")
    fmt.println("error")    //使用未导出的 方法 println错误
}
