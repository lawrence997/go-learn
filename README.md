# go-learn

这个仓库用于系统化学习 Go，目标不是只看懂示例，而是通过反复手写和项目实践，把 Go 真正用起来。

## 学习主线

当前采用这条路线：

1. 先快速过 `A Tour of Go`，建立语法和标准库的基础认识
2. 再读《The Go Programming Language》，边读边手写小例子
3. 选择一个 7 天左右的项目完整跟练
4. 跟完后独立重写一个简化版，检验是否真正掌握
5. 最后补工程基础：`testing`、`go mod`、`context`、goroutine 泄漏、`pprof`

核心原则只有一条：`写两遍`，不要只看一遍、抄一遍。

## 仓库结构

- [`basics/`](D:/codex-workspace/go-learn/basics)  
  `A Tour of Go` 跟练代码和基础语法练习
- [`flowcontrol/`](D:/codex-workspace/go-learn/flowcontrol)  
  `A Tour of Go` 中流程控制相关练习
- [`book/`](D:/codex-workspace/go-learn/book)  
  《The Go Programming Language》读书练习和手写示例
- [`projects/`](D:/codex-workspace/go-learn/projects)  
  项目跟练、功能改动、简化版重写
- [`docs/`](D:/codex-workspace/go-learn/docs)  
  学习路线、练习规则、阶段文档

## 练习要求

- 每学完一个主题，至少自己关闭教程重写一次
- 每次项目跟练，除了照着做，还要主动加一个小功能
- 代码不追求花哨，优先写清楚、写完整、写可运行
- 遇到卡点优先回到最小例子验证，而不是继续堆复杂度

## 当前阶段

当前已安排：

- `A Tour of Go` 的目录说明

接下来建议按顺序推进：

- 补 `book/` 的章节练习
- 选一个适合 7 天跟完的小项目放进 `projects/`
- 按天记录练习目标和改动
- 在 `docs/` 中维护方法和阶段总结
