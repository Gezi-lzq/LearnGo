// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 使用strings包的Join函数
	fmt.Println(strings.Join(os.Args[1:], " "))
	// 若关心输出格式，只想看看输出值
	// fmt.Println(os.Args[1:])
}
