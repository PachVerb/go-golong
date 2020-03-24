/*
    定义结构体变量, 该数据类型, 可以存储不同类型数据.
    区别数组只能存储同一类型数据。

*/

package main

import (
    "fmt"
)

//定义结构体
type Book struct {
    //声明成员字段
    bookName string
    author string
    price int
    bookId int
    
}


func main() {
  //初始化结构体
  book1 := Book{"深入node.js","hel",56,3}
  
  fmt.Println(Book{"vue.js", "刘博文", 79, 1})

  //使用key->value
  fmt.Println(Book{bookName: "目送", author: "龙应台", price:45, bookId:2})

  //或略成员字段，该字段为0 或空.(零值机制)
  //   fmt.Println(Book{"Tcp/Ip详解","key jack", 45})

  //访问 结构体成员, "."语法
  fmt.Println(book1.bookName)

  //定义结构体指针变量
  var struct_pointer *Book;

  struct_pointer = &book1

  fmt.Printf("bookName: %s\n", struct_pointer.bookName)
  fmt.Printf("bookId: %d\n", struct_pointer.bookId)

}