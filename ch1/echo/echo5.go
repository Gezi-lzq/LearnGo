// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
package main

import (
	"fmt"
	"os"
)

func main() {
	// 每次循环迭代，range产生一对值；索引以及在该索引处的元素值
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
}
