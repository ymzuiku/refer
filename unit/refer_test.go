package unit

import (
	"fmt"
	"testing"

	"github.com/ymzuiku/refer"
)

type People struct {
	Name string
	Age  int
	Age2 int64
}

func (p People) Say(a string, b int) int {
	fmt.Printf("base: %v, %v ", p.Name, p.Age)
	fmt.Printf("base: %v, %v, %v", p.Name, p.Age, p.Age2)
	fmt.Printf("say: %v, %v\n", a, b)
	return b * b
}

type NotMethod struct {
	Name string
	Age  int
	Age2 int64
}

func TestEmptyMethodSet(t *testing.T) {
	fields := refer.Fields(&NotMethod{})
	fields["Name"].SetString("dog")
	fields["Age"].SetInt(20)
	fields["Age2"].SetInt(80)

	methods := refer.Methods(People{})

	methods["Say"].Call(refer.Args("hello", 5))
}

func TestAllSet(t *testing.T) {

	fields := refer.Fields(&People{})
	fields["Name"].SetString("dog")
	fields["Age"].SetInt(20)
	fields["Age2"].SetInt(80)

	methods := refer.Methods(People{})

	out := methods["Say"].Call(refer.Args("hello", 5))

	if out[0].Interface() != 25 {
		t.Error("say error")
	}

	fiels := refer.FieldsList(&People{})
	if fiels[0] != "Name" || fiels[1] != "Age" || fiels[2] != "Age2" {
		t.Error("get field error")
	}

	m := refer.MethodsList(&People{})
	if m[0] != "Say" {
		t.Error("get method error", m)
	}
}

func TestSetAndLoad(t *testing.T) {
	target := Target{Name: "dog", Age: 10}
	fields := refer.Fields(&target)
	fields["Name"].SetString("dog2")
	if fields["Name"].Interface() != "dog2" {
		t.Error("base.Name need eq 'dog2'", fields["Name"].Interface())
	}

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
	return b.Name
}

func TestCopy(t *testing.T) {
	target := Target{Name: "dog", Age: 10}
	var base Base
	refer.Copy(&base, target)
	if base.Name != "dog" {
		t.Error("base.Name need eq 'dog'")
	}
	if base.Age != 10 {
		t.Error("bas.Agee need eq 10")
	}
	base.Say()
}

func TestCall(t *testing.T) {
	base := Base{Name: "dog"}
	values := refer.Call(&base, "Say")
	if values[0].Interface() != "dog" {
		t.Error("call error")
	}
}

func TestGetAndSet(t *testing.T) {
	base := Base{Name: "dog"}
	val := refer.Get(base, "Name")
	if val.Interface() != "dog" {
		t.Error("call error")
	}
	refer.Set(&base, "Name", "cat")
	if refer.Get(base, "Name").Interface() != "cat" {
		t.Error("call error")
	}
}
