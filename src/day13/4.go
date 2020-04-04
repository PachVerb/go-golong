/*
	字符串操作.
*/

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

//该实列实现对字符串的更新.
//a/b/c.go => c
var a = "a/b/c.go"
var b = "b/c/"
var c = "c.d.go"

//硬编码实现
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//使用string.lasIndex
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
func main() {
	//method 1
	fmt.Println(basename(a)) // c 去掉后缀 .go
	fmt.Println(basename(b)) // " "
	fmt.Println(basename(c)) //去掉.go, 结果 c.d

	//method2
	fmt.Println(basename2(a)) //c
	fmt.Println(basename2(b)) // " "
	fmt.Println(basename2(c)) //c.d

	//使用path/filepath对路径操作
	dir, file := filepath.Split(a)
	fmt.Println(dir, file)
}
