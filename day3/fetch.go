//模拟curl 发出http请求
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

func main() {
    var curl string
    for _, url := range os.Args[1:] {
        if strings.HasPrefix(url, "http") {
             curl = url + "http"
        } else {
            os.Exit(1)
        }
        resp, err := http.Get(curl)        //创建http请求
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)            //错误处理函数，发生错误时，自动终止进程
        }
        b, err := ioutil.ReadAll(resp.Body)         //获取服务器响应数据流
        resp.Body.Close()                           //关闭服务器数据流,防止资源泄漏.
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }
}