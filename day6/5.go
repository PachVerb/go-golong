/*
    赋值，= 操作数实现赋值操作
*/

package main

import "fmt"

type Stu struct {
    id int
    name string
    love string
    get  string
    your string
}
//数组
var sky =  [...]string{"23", "24", "40", "13"}
func main() {
     var a = "hello world" //普通变量赋值
     b := new(int)
     *b = 2 //指针赋值
     //结构体赋值
     var stuS = Stu{1, "jack", "ball", "game", "i love a girl"}    //声明结构体类型变量
     stuS.name="hello"
     //数组赋值
     sky[1] = "20"
     fmt.Println(a, *b, sky, stuS)

     //++语句测试
     v := 1
     v++
     fmt.Printf("test-++:%d\n", v)
}

