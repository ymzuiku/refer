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
  Age int
}

func main() {
  ref := refer.New(&People{})
  ref.F("Name").SetString("dog")
  ref.F("Age").SetInt(20)
  fmt.Println(ref.Target)
}
```
