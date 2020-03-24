/*
      求两个数最大公约数
*/
package main

import "fmt"

//每次循环交换两个位置值，对两个整数求模并交换值。
func gcd(x, y int) int {
      for y != 0 {
        x, y = y, x %y
      }
      return x
}
func main() {
    fmt.Println(gcd(12, 18))
}