# Refer

Easy use reflect, cache FieldNames

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
