/*
	类型声明语句
	type 类型名称 底层类型
	类型声明，即使有相同底层类型，但数据类型不同，也是不能进行表示式运算和比较运算的。

	对于每一个类型T,都有对应类型T(x)的转换操作
	如果T是指针类型，可能会需要用小括弧包装T，比如(*int)(0)）。
	只有当两个类型的底层基础类型相同时，才允许这种转型操作，
	或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。
	如果x是可以赋值给T类型的值，那么x必然也可以被转为T类型，但是一般没有这个必要。
*/
//该包用于换算华氏温度和摄氏温度
package tempconv //设置包名

import "fmt"

type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

//格式化操作
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
