// 题目：Functions continued（函数续）
// 中文说明：学习多个相邻参数类型相同时的简写方式。
// 官方来源：https://go.dev/tour/basics/5
// 题目列表：https://go.dev/tour/list
// 练习建议：
// 1. 把 func add(x int, y int) int 改成 func add(x, y int) int。
// 2. 对比两种写法的可读性。
// 3. 自己再写一个接受三个参数的函数。

package main

import (
	"fmt"
	"math/cmplx"
)

func add(a,y int)(int,int){
	return a,y
}

func add02(a,y int)(a1,y1 int){
    a1 = a + y
	y1 = 2 + 3
	return 
}

var(
	b,bol,java bool
	J1 ,  J2 =  "ccc" , "dddd"
	d1 , d2 = "c1","2"
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var i int
	
	// TODO: 在这里写你的练习代码
	a,y := add(1,2)
	fmt.Println(a,y)
	a,y = add02(1,2)
	fmt.Println(a,y)
	fmt.Println(i,b,bol,java)
	fmt.Println(d1,d2)
	fmt.Printf("开始展示： %T %v \n",J1,J2)
	
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

