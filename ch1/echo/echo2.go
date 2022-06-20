// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用一条短变量声明来声明并初始化s和seps
	// 只能用在函数内部，而不能用于包变量
	s, sep := "", ""
	// 每次循环迭代，range产生一对值；索引以及在该索引处的元素值
	for _, arg := range os.Args[1:] {
		// 空标识符_可用于在任何语法需要变量名但程序逻辑不需要的时候
		//（如：在循环里）丢弃不需要的循环索引，并保留元素值
		s += sep + arg
		// +=连接原字符串、空格和下个参数，产生新字符串，并把它赋值给s。
		// s原来的内容已经不再使用，将在适当时机对它进行垃圾回收。
		// 如果连接涉及的数据量很大，这种方式代价高昂
		sep = " "
	}
	fmt.Println(s)
}
