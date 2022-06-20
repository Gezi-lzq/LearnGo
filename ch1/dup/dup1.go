// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打印标准输入中多次出现的行，以重复次数开头
// 灵感来自于Unix的uniq命令，其寻找相邻的重复行
func main() {
	// 内置函数make创建空map (key:string, value:int)
	counts := make(map[string]int)
	// 使用短变量声明创建bufio.Scanner类型的变量input
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // 每次调用input.Scan(),即读入下一行,并移除行末的换行符
		counts[input.Text()]++ // 读取的内容可以调用input.Text()得到
	}
	// NOTE: ignoring potential errors from input.Err()
	// map的迭代顺序随机
	for line, n := range counts {
		if n > 1 {
			// 格式化输出
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
