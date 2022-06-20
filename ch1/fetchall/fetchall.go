// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

/*
	fetchall的特别之处在于它会同时去获取所有的URL
	所以这个程序的总执行时间不会超过执行时间最长的那一个任务
	前面的fetch程序执行时间则是所有任务执行时间之和。
	fetchall程序只会打印获取的内容大小和经过的时间，不会像之前那样打印获取的内容。
*/
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// 用make函数创建了一个传递string类型参数的channel
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 对每一个命令行参数，我们都用go这个关键字来创建一个goroutine
		// 并且让函数在这个goroutine异步执行http.Get方法
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		// 主函数负责接收这些值（<-ch）
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	// io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中
	// (译注：可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据)
	// 因为我们需要这个方法返回的字节数，但是又不想要其内容
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	// 每一个fetch函数在执行时都会往channel里发送一个值（ch <- expression）
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
	// 当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，
	// 直到另一个goroutine从这个channel里接收或者写入值，这样两个goroutine才会继续执行channel操作之后的逻辑。
}
