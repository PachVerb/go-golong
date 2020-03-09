package main

import (
    "fmt"      //已包名导入
    "math/rand"           //路径导入包，注意：这里导入包的路径中最后一个元素，因与源码开始名字一致
    "time"
)

//说明: rand.Intn伪随机函数，默认资源是单一值.  Intn arguments:[number] 传入参数表示生成随机数范围
//需要调用 rand.seed传入变化值,初始化资源库
func main(){
    fmt.Println("生成数字", rand.Intn(10));      //print number 1

    fmt.Println("print time:", time.Now())   //返回当前时间
    fmt.Println("print random time:", time.Now().UnixNano())   //毫秒时间
    
    rand.Seed(time.Now().UnixNano())   //初始化种子
    fmt.Println("a number of 1-1000", rand.Intn(1000))
}