// 题目：Exported names（导出名）
// 中文说明：学习 Go 中首字母大写表示导出，小写表示包内可见。
// 官方来源：https://go.dev/tour/basics/3
// 题目列表：https://go.dev/tour/list
// 练习建议：
// 1. 观察 math.Pi 和 math.pi 的区别。
// 2. 理解“能否跨包访问”与“名字是否导出”的关系。
// 3. 自己定义一个导出函数和一个未导出函数做对比。

package main

import (
	"fmt"
	"math"
)

/*
小写函数是私有的，不能被其他包访问
大写函数是公共的，可以被其他包访问
*/
func main() {
	// TODO: 在这里写你的练习代码
	fmt.Print(math.pi)
}
