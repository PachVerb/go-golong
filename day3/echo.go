package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
//         s += sep + arg
//         sep = " "
        fmt.Println(_)
        fmt.Println(arg)
    }
    fmt.Println(s)
    fmt.Println(os.Args[0])
//     fmt.Println(os.Args[2])
}