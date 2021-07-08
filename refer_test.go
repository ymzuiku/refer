package refer

import (
	"fmt"
	"testing"
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

type NotMethod struct {
	Name string
	Age  int
	Age2 int64
}

func TestEmptyMethodSet(t *testing.T) {
	ref := New(NotMethod{})
	ref.F("Name").SetString("dog")
	ref.F("Age").SetInt(20)
	ref.F("Age2").SetInt(80)
	fmt.Println(ref.Interface())
	if _, err := ref.Call("Say", "hello", 5); !IsErrCallEmptyMethods(err) {
		t.Error(err)
	}
}

func TestAllSet(t *testing.T) {
	ref := New(People{})
	ref.F("Name").SetString("dog")
	ref.F("Age").SetInt(20)
	ref.F("Age2").SetInt(80)
	fmt.Println(ref.Interface())
	if _, err := ref.Call("Say", "hello", 5); err != nil {
		t.Error(err)
	}
	list := ref.GetFieldNames()
	if list[0] != "Name" || list[1] != "Age" || list[2] != "Age2" {
		t.Error("get field error")
	}

	if _, err := ref.Call("No_Say", "hello", 5); !IsErrCallNotHaveMethod(err) {
		t.Error(err)
	}
}
