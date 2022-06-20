package main

import (
	"fmt"
	"os"
)

func main() {
	// var声明定义了两个string类型的变量s和sep。
	// 变量会在声明时直接初始化。
	// 如果变量没有显式初始化，则被隐式地赋予其类型的零值（zero value）
	// 数值类型是0，字符串类型是空字符串""
	var s, sep string
	// os.Args变量是一个字符串（string）的切片（slice）
	// 符号:=是短变量声明
	for i := 1; i < len(os.Args); i++ { // 二次加工（quadratic process）拼凑参数
		s += sep + os.Args[i] // 当参数数量庞大时，开销会很大
		sep = " "
	}
	fmt.Println(s)
}
