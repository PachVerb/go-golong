/*
	math包，sqrt
*/
package main

import (
	"fmt"
	"math"
)

var i float64 = 4

func main() {
	a, b := 2, 2
	//sqrt
	fmt.Printf("%g\n", math.Sqrt(i)) //2
	//Hypot
	fmt.Println(math.Hypot(float64(a), float64(b))) //返回两数平方和二次方根
	//sin
	fmt.Println(math.Sin(float64(math.Pi / 6)))
}
