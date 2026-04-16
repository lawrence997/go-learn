// 题目：Packages（包）
// 中文说明：学习 Go 程序由 package 组成，入口通常是 package main。
// 官方来源：https://go.dev/tour/basics/1
// 题目列表：https://go.dev/tour/list
// 练习建议：
// 1. 先自己写出 package 和 import。
// 2. 观察 main 函数为什么是程序入口。
// 3. 尝试引入 fmt、math/rand 等包做最小示例。

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 程序入口
	fmt.Print("my favorite number is ", rand.Intn(10), "\n")
}
