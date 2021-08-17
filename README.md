# Refer

Easy use reflect, only cache FieldNames

## Installation

```bash
go get github.com/ymzuiku/refer
```

## Use fields

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
	fields := refer.Fields(&People{})
	fields["Name"].SetString("dog")
	fields["Age"].SetInt(20)
	fields["Age2"].SetInt(80)
}
```

## Use copy

```go
import (
  "fmt"
  "github.com/ymzuiku/refer"
)

type Base struct {
	Name string
}

type A struct {
	Name string
}

type B struct {
	Name string
}

func say(v Base) {
	var base Base
	refer.Copy(&base, v)
	fmt.Printf("base: %v, %v ", b.Name)
}

func main() {
	say(A{Name:"Apple"})
	say(A{Name:"Banana"})
}
```

## Use call

```go
import (
  "fmt"
  "github.com/ymzuiku/refer"
)

type Base struct {
	Name string
}

type (b Base) Say(hello string){
	fmt.Println(hello, b.Name)
}

func main() {
	refer.Call(Base{Name:"dog"}, "Say", "Hi, ")
}
```

## Benchmark

```
BenchmarkReflectNew-8                   29918192                36.44 ns/op           64 B/op        1 allocs/op
BenchmarkReflectSet-8                   62631976                55.63 ns/op            0 B/op        0 allocs/op
BenchmarkReflectNameSet-8                4362084               275.6 ns/op            32 B/op        4 allocs/op
BenchmarkReflectCacheNameSet-8          19608376                76.95 ns/op            0 B/op        0 allocs/op
BenchmarkReferNameSet-8                 19948922                76.33 ns/op            0 B/op        0 allocs/op
BenchmarkReflectMethodIndexCall-8        3480391               342.8 ns/op            95 B/op        5 allocs/op
BenchmarkReflectMethodNameCall-8         1924236               624.4 ns/op           311 B/op        9 allocs/op
BenchmarkReferCall-8                     3020980               391.2 ns/op           143 B/op        6 allocs/op
BenchmarkReferCopy-8                     2103394               571.3 ns/op           896 B/op       11 allocs/op
```
