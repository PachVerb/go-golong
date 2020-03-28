/*
	该包用于bool类型转 1， 0
	原由: bool类型不能隐式转换. go各算数运算，比较运算。强制要求类型匹配.
*/
package conv

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
