package main

//包组合形式导入;
import (
    "fmt"
    "math"
)

//分组形式导入
import "time"
import "math/rand"

func main(){
    //获取一个数的平方根
    fmt.Println("get number", math.Sqrt(4))
    
    //获取随机数
    rand.Seed(time.Now().UnixNano())
    fmt.Println("a number of 1-40:", rand.Intn(40))
}