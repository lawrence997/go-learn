# Go 标准库速查表

这份文档的目标不是让你死背全部标准库，而是建立一个稳定的索引：

- 遇到问题时，知道先想到哪个包
- 知道每个高频包最常见的用途
- 写练习或项目时，可以快速翻看

---

## 学习建议

- 不用背全
- 先掌握高频包
- 按“需求 -> 包名”来记忆
- 写代码时配合官方文档和 IDE 提示

一个更实用的目标是：

- 想打印内容，先想到 `fmt`
- 想处理字符串，先想到 `strings`
- 想做 JSON，先想到 `encoding/json`
- 想写 HTTP 服务，先想到 `net/http`
- 想处理并发，先想到 `sync` 和 `context`

---

## 初学者优先掌握

如果只先学一批，建议优先熟悉下面这些包：

- `fmt`
- `os`
- `io`
- `bufio`
- `strings`
- `strconv`
- `time`
- `errors`
- `sort`
- `encoding/json`
- `net/http`
- `sync`
- `context`
- `testing`

---

## 1. 基础输入输出

### `fmt`

用途：格式化输入输出。

常见函数：

- `fmt.Println()`
- `fmt.Printf()`
- `fmt.Sprintf()`
- `fmt.Scan()`

示例：

```go
package main

import "fmt"

func main() {
	name := "Tom"
	age := 18
	fmt.Printf("name=%s age=%d\n", name, age)
}
```

### `bufio`

用途：带缓冲的读写，常用于标准输入或文件读取。

常见函数：

- `bufio.NewReader()`
- `bufio.NewWriter()`
- `ReadString('\n')`

示例：

```go
reader := bufio.NewReader(os.Stdin)
line, _ := reader.ReadString('\n')
```

### `io`

用途：抽象读写接口和通用读写工具。

常见内容：

- `io.Reader`
- `io.Writer`
- `io.Copy()`
- `io.ReadAll()`

示例：

```go
data, _ := io.ReadAll(resp.Body)
```

---

## 2. 文件、系统、路径

### `os`

用途：文件、进程、环境变量、命令行参数。

常见函数或内容：

- `os.Open()`
- `os.Create()`
- `os.ReadFile()`
- `os.WriteFile()`
- `os.Args`
- `os.Getenv()`

示例：

```go
content, _ := os.ReadFile("a.txt")
fmt.Println(string(content))
```

### `path/filepath`

用途：处理本地文件路径，尽量避免手写斜杠拼接。

常见函数：

- `filepath.Join()`
- `filepath.Base()`
- `filepath.Dir()`
- `filepath.Ext()`
- `filepath.Walk()`

示例：

```go
p := filepath.Join("data", "user", "a.txt")
fmt.Println(p)
```

### `io/fs`

用途：统一文件系统接口，常见于遍历文件、抽象文件系统、嵌入资源。

常见内容：

- `fs.FS`
- `fs.WalkDir()`

---

## 3. 字符串与文本处理

### `strings`

用途：字符串切分、拼接、查找、替换。

常见函数：

- `strings.Contains()`
- `strings.Split()`
- `strings.Join()`
- `strings.TrimSpace()`
- `strings.ReplaceAll()`
- `strings.HasPrefix()`
- `strings.HasSuffix()`

示例：

```go
s := "a,b,c"
parts := strings.Split(s, ",")
fmt.Println(parts)
```

### `strconv`

用途：字符串和数字、布尔值之间的转换。

常见函数：

- `strconv.Atoi()`
- `strconv.Itoa()`
- `strconv.ParseFloat()`
- `strconv.ParseBool()`
- `strconv.FormatInt()`

示例：

```go
n, _ := strconv.Atoi("123")
fmt.Println(n + 1)
```

### `bytes`

用途：处理 `[]byte`，很多接口、网络和文件操作里都会遇到。

常见函数：

- `bytes.Contains()`
- `bytes.Split()`
- `bytes.NewBuffer()`
- `bytes.TrimSpace()`

### `unicode`

用途：判断字符类型。

常见函数：

- `unicode.IsLetter()`
- `unicode.IsDigit()`
- `unicode.IsSpace()`

### `regexp`

用途：正则表达式匹配和提取。

常见函数：

- `regexp.MustCompile()`
- `MatchString()`
- `FindString()`
- `FindStringSubmatch()`

示例：

```go
re := regexp.MustCompile(`\d+`)
fmt.Println(re.FindString("abc123xyz"))
```

---

## 4. 时间、数学、随机数

### `time`

用途：时间日期、定时、超时控制。

常见函数：

- `time.Now()`
- `time.Sleep()`
- `time.Parse()`
- `time.Format()`
- `time.Since()`
- `time.After()`

示例：

```go
now := time.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

注意：Go 的时间格式不是 `YYYY-MM-DD`，而是固定参考时间：

`2006-01-02 15:04:05`

### `math`

用途：数学计算。

常见函数：

- `math.Abs()`
- `math.Sqrt()`
- `math.Pow()`
- `math.Ceil()`
- `math.Floor()`

### `math/rand`

用途：伪随机数。

常见函数：

- `rand.Intn()`
- `rand.Float64()`
- `rand.Shuffle()`

示例：

```go
fmt.Println(rand.Intn(100))
```

说明：如果是密码学或安全场景，不要用 `math/rand`，应考虑 `crypto/rand`。

---

## 5. 数据编码与解码

### `encoding/json`

用途：JSON 编码与解码，高频包。

常见函数：

- `json.Marshal()`
- `json.Unmarshal()`
- `json.NewEncoder()`
- `json.NewDecoder()`

示例：

```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

u := User{Name: "Tom", Age: 18}
b, _ := json.Marshal(u)
fmt.Println(string(b))
```

### `encoding/base64`

用途：Base64 编解码。

### `encoding/csv`

用途：CSV 文件读写。

### `encoding/xml`

用途：XML 编解码。

---

## 6. 错误、日志、运行时

### `errors`

用途：创建和判断错误。

常见函数：

- `errors.New()`
- `errors.Is()`
- `errors.As()`

示例：

```go
err := errors.New("something wrong")
fmt.Println(err)
```

### `log`

用途：基础日志输出。

常见函数：

- `log.Println()`
- `log.Printf()`
- `log.Fatalf()`
- `log.SetFlags()`

### `runtime`

用途：查看运行时环境信息。

常见内容：

- `runtime.GOOS`
- `runtime.GOARCH`
- `runtime.NumCPU()`

---

## 7. 集合、排序、反射

### `sort`

用途：排序。

常见函数：

- `sort.Ints()`
- `sort.Strings()`
- `sort.Slice()`

示例：

```go
nums := []int{5, 2, 8, 1}
sort.Ints(nums)
fmt.Println(nums)
```

### `slices`

用途：切片操作工具，Go 新版本里很实用。

常见函数：

- `slices.Sort()`
- `slices.Contains()`
- `slices.Index()`
- `slices.Clone()`

### `maps`

用途：map 工具函数。

常见函数：

- `maps.Clone()`
- `maps.Equal()`

### `reflect`

用途：反射，常见于框架、序列化、通用工具库。

入门阶段只需要先知道它的存在和用途，不用急着深入。

---

## 8. 网络开发

### `net/http`

用途：HTTP 客户端和服务端开发核心包。

常见函数：

- `http.Get()`
- `http.Post()`
- `http.HandleFunc()`
- `http.ListenAndServe()`

服务端示例：

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
})
http.ListenAndServe(":8080", nil)
```

客户端示例：

```go
resp, _ := http.Get("https://example.com")
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)
fmt.Println(string(body))
```

### `net/url`

用途：URL 解析和查询参数处理。

常见内容：

- `url.Parse()`
- `url.Values`

### `net`

用途：更底层的网络编程，比如 TCP、UDP。

---

## 9. 并发相关

### `sync`

用途：并发同步。

常见类型：

- `sync.Mutex`
- `sync.RWMutex`
- `sync.WaitGroup`
- `sync.Once`

示例：

```go
var wg sync.WaitGroup
wg.Add(1)

go func() {
	defer wg.Done()
	fmt.Println("working")
}()

wg.Wait()
```

### `sync/atomic`

用途：原子操作，常见于计数器、状态位。

### `context`

用途：超时、取消、请求链路传递。

常见函数：

- `context.Background()`
- `context.WithCancel()`
- `context.WithTimeout()`
- `context.WithValue()`

示例：

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

---

## 10. 测试相关

### `testing`

用途：单元测试、基准测试。

常见形式：

- `func TestXxx(t *testing.T)`
- `func BenchmarkXxx(b *testing.B)`

示例：

```go
func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("want 3")
	}
}
```

### `net/http/httptest`

用途：测试 HTTP 服务。

---

## 11. 命令行与配置

### `flag`

用途：解析命令行参数。

示例：

```go
name := flag.String("name", "guest", "user name")
flag.Parse()
fmt.Println(*name)
```

### `os`

也常用于直接读取命令行参数：

- `os.Args`

---

## 12. 模板

### `text/template`

用途：生成普通文本。

### `html/template`

用途：生成 HTML，并自动处理一部分转义问题。

示例：

```go
tmpl, _ := template.New("x").Parse("<h1>{{.}}</h1>")
tmpl.Execute(os.Stdout, "Hello")
```

---

## 13. 加密与安全

### `crypto/rand`

用途：安全随机数。

### `crypto/sha256`

用途：SHA256 哈希。

### `crypto/md5`

用途：MD5 哈希。

注意：很多安全场景不推荐继续使用 MD5。

---

## 14. 常见高频组合

你以后写代码时，经常会把这些包组合起来用：

- 读文件：`os` + `io`
- 处理文本：`strings` + `strconv`
- 做 HTTP 接口：`net/http` + `encoding/json`
- 控制超时：`time` + `context`
- 并发任务：`go` + `sync`
- 测试接口：`testing` + `net/http/httptest`

---

## 15. 一句话速记版

- `fmt`：打印
- `os`：系统和文件
- `io`：读写抽象
- `bufio`：缓冲读写
- `strings`：字符串处理
- `strconv`：字符串转换
- `time`：时间
- `math`：数学
- `sort`：排序
- `errors`：错误
- `log`：日志
- `encoding/json`：JSON
- `net/http`：HTTP
- `sync`：并发同步
- `context`：超时取消
- `testing`：测试

---

## 16. 记忆方法

不要按“包名单词表”硬背，尽量按下面的映射来记：

- 打印输出：`fmt`
- 读输入：`bufio`
- 文件读写：`os`
- 字符串切分：`strings`
- 字符串转数字：`strconv`
- 时间处理：`time`
- JSON：`encoding/json`
- HTTP：`net/http`
- 排序：`sort`
- 并发等待：`sync`
- 错误处理：`errors`
- 命令行参数：`flag`

这类映射记熟之后，标准库就会慢慢从“背知识点”变成“解决问题的工具箱”。
