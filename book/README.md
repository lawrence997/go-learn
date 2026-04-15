# book

这个目录用于存放《The Go Programming Language》的读书练习和手写示例。

## 目标

- 用系统阅读补足 Go 语言机制层面的理解
- 不只知道语法怎么写，还要知道为什么这样设计
- 为后续项目练习打基础，尤其是 `package`、接口、并发、错误处理

## 建议做法

- 按章节建文件或子目录
- 每章至少保留一个自己手写的小例子
- 对关键概念写出最小可运行代码，而不是只摘抄书中结论
- 遇到抽象内容时，优先写实验代码验证

## 可以重点练习的主题

- 包与可见性
- 指针与方法
- 接口与组合
- goroutine 与 channel
- 错误处理
- I/O 与标准库使用

## 文件建议

可以按类似方式组织：

- `ch01-hello.go`
- `ch02-package.go`
- `ch03-functions.go`
- `ch04-slices-maps.go`
- `ch05-methods-interfaces.go`
- `ch06-goroutines-channels.go`

重点不是文件名，而是确保每一部分都有自己的手写版本。
