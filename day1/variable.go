//变量
//方式：var声明变量。变量名在前，类型在后. 同函数参数列表一样。
//变量声明可以出现在导入包和函级别
//变量声明若果初始化，可以省略类型。go会根据值判断类型
package main
                            
import "fmt"

var x, y bool //声明一组相同类型变量.

//这里有个问题,在未初始化时，声明一组不同类型变量报错。

func main() {
    var num int = 2
    var a, b, c = 1, true, "xialuoke"      //初始化时，申明一组不同类型变量。

    fmt.Println( x, y, num, a, b, c)        //false
}