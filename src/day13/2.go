/*
	利用go的utf-8编码有点，字符串操作不需要转码操作

*/
package main

import "fmt"

//测试一个字串是否是另一个字串前缀
var a, b = "hello,", "hello, world"

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

//测试是否为后缀
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

//测试包含子串测试：
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
func main() {
	//测试前缀
	fmt.Println(HasPrefix(a, b)) //false
	fmt.Println(HasPrefix(b, a)) //true
	fmt.Println(HasSuffix(a, b)) //false
	fmt.Println(HasSuffix(b, a)) //false
	fmt.Println(a[:6])

	//字串测试
	fmt.Println(Contains(a, b)) //false
}
