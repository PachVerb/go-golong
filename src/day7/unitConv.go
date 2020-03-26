package main

import (
	"./unit_conv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//range 产生产生索引和值
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err) //os.stderr标准错误日志输出 ,
			os.Exit(1)                              //发生错误，退出程序
		}

		//输入单位相互转换
		k := conv.Kilogram(t)
		g := conv.Gram(t)

		fmt.Printf("%s = %s, %s = %s\n", k, conv.Gconv(k), g, conv.Kconv(g))
		//fmt.Println(arg)	//循环拿出每一个参数值
	}
	//fmt.Println(os.Args[1]) 	//os.Args获取命令行输入命令。os.Args[0]获取命令本身体. os.Args[m:n]数组切片
	//fmt.Println(os.Args[1:])
}
