# Refer

Easy use reflect, only cache FieldNames

## Installation

```bash
go get github.com/ymzuiku/refer
```

# Usage

```go
import (
  "fmt"
  "github.com/ymzuiku/refer"
)
type People struct {
	Name string
	Age  int
	Age2 int64
}

func (p People) Say(a string, b int) {
	fmt.Printf("base: %v, %v ", p.Name, p.Age)
	fmt.Printf("base: %v, %v, %v", p.Name, p.Age, p.Age2)
	fmt.Printf("say: %v, %v\n", a, b)
}
func main() {
	ref := New(People{})
	ref.F("Name").SetString("dog")
	ref.F("Age").SetInt(20)
	ref.F("Age2").SetInt(80)
	if _, err := ref.Call("Say", "hello", 5); err != nil {
    t.Error(err)
	}
  fmt.Println(ref.Interface())
}
```

## Benchmark

| Benchmark Name                  | 说明                                   |
| ------------------------------- | -------------------------------------- |
| BenchmarkNew                    | 直接 new                               |
| BenchmarkReflectNew             | 使用 reflect new                       |
| BenchmarkSet                    | 直接赋值 field                         |
| BenchmarkReflectSet             | reflect index 设置 field               |
| BenchmarkReflectNameSet         | 不缓存 fieldName 设置 field            |
| BenchmarkReflectCacheNameSet    | 缓存 fieldName 设置 field              |
| BenchmarkReferNameSet           | 通过 refer 的 Field[name] 设置 field   |
| BenchmarkCall                   | 直接执行函数                           |
| BenchmarkReflectMethodIndexCall | 使用 reflect method index 执行函数     |
| BenchmarkReflectMethodNameCall  | 使用 reflect methodName index 执行函数 |
| BenchmarkReferCall              | 使用 refer.Call(name) 执行函数         |

```
BenchmarkNew-8                          57849320                20.90 ns/op           64 B/op          1 allocs/op
BenchmarkReflectNew-8                   36778461                33.48 ns/op           64 B/op          1 allocs/op
BenchmarkSet-8                          1000000000               0.3240 ns/op          0 B/op          0 allocs/op
BenchmarkReflectSet-8                   56485318                21.36 ns/op            0 B/op          0 allocs/op
BenchmarkReflectNameSet-8                4283246               283.9 ns/op            32 B/op          4 allocs/op
BenchmarkReflectCacheNameSet-8          37124451                31.67 ns/op            0 B/op          0 allocs/op
BenchmarkReferNameSet-8                 38366032                31.78 ns/op            0 B/op          0 allocs/op
BenchmarkReferNameUseGet-8              38153932                31.55 ns/op            0 B/op          0 allocs/op
BenchmarkCall-8                         1000000000               0.3197 ns/op          0 B/op          0 allocs/op
BenchmarkReflectMethodIndexCall-8        4507449               267.9 ns/op           119 B/op          4 allocs/op
BenchmarkReflectMethodNameCall-8         2100422               570.3 ns/op           335 B/op          8 allocs/op
BenchmarkReferCall-8                     3743725               319.8 ns/op           167 B/op          5 allocs/op
```
