// Server1 is a minimal "echo" server.
package main

/*
	这个服务器的功能是返回当前用户正在访问的URL。
	比如用户访问的是 http://localhost:8000/hello ，那么响应是URL.Path = "hello"。
*/
import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//   main函数将所有发送到/路径下的请求和handler函数关联起来
	//  /开头的请求其实就是所有发送到当前站点上的请求，服务监听8000端口
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	// ParseForm被嵌套在了if语句中
	// 允许这样简单的语句结果作为局部的变量声明出现在if语句的最前面
	// 用if和ParseForm结合可以让代码更加简单，并且可以限制err这个变量的作用域
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
