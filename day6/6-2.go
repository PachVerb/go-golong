/*
    求斐波拉契数列第n个数 
*/

package main

import "fmt"

func numList(n int) int {
    x, y := 1, 2

    for i := 0; i < n; i++  {
        x, y = y, x + y
    }
    return x
}
func main() {
    //前5项，1, 2, 3, 5, 8, 13.每项从0计数
    fmt.Println(numList(4))
}