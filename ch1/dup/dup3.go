package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// 其读取指定文件的全部内容
		// ReadFile函数返回一个字节切片（byte slice）
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// strings.Split函数把字符串分割成子串的切片
		// 字节切片（byte slice）转换为string，才能用strings.Split分割
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
