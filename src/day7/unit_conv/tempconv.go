package conv

import "fmt"

type Kilogram float64
type Gram float64

//导出变量名大写
const (
	//K
	A Kilogram = 2
	B Kilogram = 4
	C Kilogram = 6

	//G
	Ga Gram = 100
	Gb Gram = 200
	Gc Gram = 300
)

//此处的声明语句，标明声明一个Kilogram类型的string函数。类型声明很多时候需要有string函数。目的是格式化。
func (k Kilogram) String() string {
	return fmt.Sprintf("%gKG", k)
}
func (g Gram) String() string {
	return fmt.Sprintf("%gG", g)
}
