// 练习 1.8：
// 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀。
// 你可能会用到strings.HasPrefix这个函数
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// 如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		// http.Get函数是创建HTTP请求的函数
		// 会在resp这个结构体中得到访问的请求结果
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// resp的Body字段包括一个可读的服务器响应流
		dst := os.Stdout
		_, err = io.Copy(dst, resp.Body)

		// resp.Body.Close关闭resp的Body流，防止资源泄露
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
