package main

import (
	"bytes"
	"fmt"
)

//字符串整数，每隔三个字符插入comma
var a = "12314234235"

//func addComma(s string) string {
//	for i := len(s)-1; i >=0; i-=3{
//		if len(s) < 3 {
//			break
//		}
//		s = s[:i] + "," + s[i-3:]
//	}
//	return s
//}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:] //递归调用修改s
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
func main() {
	fmt.Println(comma(a)) // 12, 314, 234, 235
	//fmt.Println(addComma(a))

	//字符串和splice转换
	s := "hello, world"
	b := []byte(s)
	s2 := string(b) //字节slice转到字符串。 目的是保证s2只读

	fmt.Println(s, b, s2)
	fmt.Printf("%T\n", s2) //string

	//bytes.包buffer
	fmt.Println(intsToString([]int{1, 2, 3})) //[1, 2, 3]

}
