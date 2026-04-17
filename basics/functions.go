// 题目：Functions（函数）
// 中文说明：学习函数参数、返回值，以及 Go 的类型写在变量名后面。
// 官方来源：https://go.dev/tour/basics/4
// 题目列表：https://go.dev/tour/list
// 练习建议：
// 1. 写一个 add(x int, y int) int。
// 2. 观察参数和返回值类型的声明位置。
// 3. 再自己写一个减法或乘法函数。
// *是获取地址，声明指针的语法。
// &是通过指针

package main

import "fmt"

func add(x int, y int) int{
	return x*y
}

func main() {
	// TODO: 在这里写你的练习代码
	fmt.Print(add(100,2))
}
