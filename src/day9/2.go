/*
	获取当前工作目录， 演示局部变量和局部变量
*/
package main

import (
	"fmt"
	"os"
)

var cwd string

/*
	程序正常运行，没有出错。其实是隐含bug的。包级别变量自始至终，都没有得到更新。由于main函数内部声明同名变量屏蔽外部变量因此获取的结果
	在局部cwd中。为避免这种情况，可以单独声明err.
*/
func main() {
	var err error
	cwd, err = os.Getwd()

	if err != nil {
		fmt.Printf("get wd failed: %v", err)
	}
	fmt.Printf("%v\n", cwd) // -->/Users/ousan/Desktop/goLearnNot/src/day9
}
