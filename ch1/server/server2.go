// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// 必须保证每次修改变量的最多只能有一个goroutine
// 这也就是代码里的mu.Lock()和mu.Unlock()调用将修改count的所有行为包在中间的目的
var mu sync.Mutex
var count int

func main() {
	// 根据请求的url不同会调用不同的函数
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
