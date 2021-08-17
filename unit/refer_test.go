package unit

import (
	"fmt"
	"testing"
	"unsafe"

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

type NotMethod struct {
	Name string
	Age  int
	Age2 int64
}

func TestEmptyMethodSet(t *testing.T) {
	ref := refer.New(NotMethod{})
	ref.F("Name").SetString("dog")
	ref.F("Age").SetInt(20)
	ref.F("Age2").SetInt(80)
	fmt.Println(ref.Interface())
	ref.Call("Say", "hello", 5)
}

func TestAllSet(t *testing.T) {

	ref := refer.New(People{})

	fmt.Println(unsafe.Sizeof(People{}))
	fmt.Println(unsafe.Sizeof(ref))

	ref.F("Name").SetString("dog")
	ref.F("Age").SetInt(20)
	ref.F("Age2").SetInt(80)
	fmt.Println(ref.Interface())
	ref.Call("Say", "hello", 5)
	fiels := ref.GetFieldNames()
	if fiels[0] != "Name" || fiels[1] != "Age" || fiels[2] != "Age2" {
		t.Error("get field error")
	}

	methods := ref.GetMethodNames()
	if methods[0] != "Say" {
		t.Error("get method error")
	}

	ref.Call("No_Say", "hello", 5)
}

type Base struct {
	Name string
	Age  int
	Haha func() string
}

func (b Base) Say() string {
	fmt.Println("base say", b.Name, b.Age)
	return b.Name
}

type Target struct {
	Name string
	Age  int
}

func (b Target) Say() string {
	fmt.Println("target say", b.Name, b.Age)
	return b.Name
}

// func TestCopy(t *testing.T) {
// 	// target := Target{Name: "dog", Age: 10}
// 	// base := Base{}

// 	target := Target{Name: "dog", Age: 10}
// 	base := Base{}
// 	refer.Copy(&base, target)
// 	fmt.Println(base)
// 	if base.Name != "dog" {
// 		t.Error("base.Name need eq 'dog'")
// 	}
// 	if base.Age != 10 {
// 		t.Error("bas.Agee need eq 10")
// 	}
// 	base.Say()
// 	t.Error("aa")
// }
