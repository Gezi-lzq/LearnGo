// 练习 1.9：
// 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// http.Get函数是创建HTTP请求的函数
		// 会在resp这个结构体中得到访问的请求结果
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// resp的Body字段包括一个可读的服务器响应流
		// ioutil.ReadAll函数从response中读取到全部内容；将其结果保存在变量b中
		b, err := ioutil.ReadAll(resp.Body)
		code := resp.StatusCode //状态码
		// resp.Body.Close关闭resp的Body流，防止资源泄露
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// Printf函数会将结果b写出到标准输出流中
		fmt.Printf("StatusCode is %d \n %s", code, b)
	}
}
