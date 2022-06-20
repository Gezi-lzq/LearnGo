// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 读取标准输入或是使用 os.Open 打开各个具名文件，并操作它们
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// 从标准输入中读取数据
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// 返回第一个值是被打开的文件(*os.File)
			// 返回的第二个值是内置error类型的值
			f, err := os.Open(arg)
			// 如果err的值不是nil，说明打开文件时出错了
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// 从文件输入中读取数据
			countLines(f, counts)
			// 调用Close关闭该文件，并释放占用的所有资源
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 函数和包级别的变量（package-level entities）可以任意顺序声明，并不影响其被调用
func countLines(f *os.File, counts map[string]int) {
	// map作为参数传递给某函数时
	// 该函数接收这个引用的一份拷贝（copy，或译为副本）
	// 被调用函数对map底层数据结构的任何修改
	// 调用者函数都可以通过持有的map引用看到
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		// 练习 1.4： 修改dup2，出现重复的行时打印文件名称
		if counts[input.Text()] > 1 {
			fmt.Println(input.Text())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
