package unit

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ymzuiku/refer"
)

type Config struct {
	Name    string `json:"server-name"`
	IP      string `json:"server-ip"`
	URL     string `json:"server-url"`
	Timeout string `json:"timeout"`
}

func (c Config) Sub(a, b int) (int, string) {
	return a - b, "done"
}

func (c Config) Sub2(a, b int) (int, string) {
	return a - b, "done"
}

func (c Config) Sub3(a, b int) (int, string) {
	return a - b, "done"
}

func (c Config) Add(a, b int) (int, string) {
	return a + b, "done"
}

func BenchmarkNew(b *testing.B) {
	var config *Config
	for i := 0; i < b.N; i++ {
		config = &Config{}
	}
	_ = config
}

func BenchmarkReflectNew(b *testing.B) {
	var config *Config
	typ := reflect.TypeOf(Config{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config, _ = reflect.New(typ).Interface().(*Config)
	}
	_ = config
}

func BenchmarkSet(b *testing.B) {
	config := &Config{}
	for i := 0; i < b.N; i++ {
		config.Name = "name"
		config.IP = "ip"
		config.URL = "url"
		config.Timeout = "timeout"
	}
}

func BenchmarkReflectSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	for i := 0; i < b.N; i++ {
		ins.Field(0).SetString("name")
		ins.Field(1).SetString("ip")
		ins.Field(2).SetString("url")
		ins.Field(3).SetString("timeout")
	}
}

func BenchmarkReflectNameSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	for i := 0; i < b.N; i++ {
		ins.FieldByName("Name").SetString("name")
		ins.FieldByName("IP").SetString("ip")
		ins.FieldByName("URL").SetString("url")
		ins.FieldByName("Timeout").SetString("timeout")
	}
}

func ReflectCache(target interface{}) map[string]reflect.Value {
	typ := reflect.TypeOf(target)
	ins := reflect.New(typ).Elem()
	fields := make(map[string]reflect.Value, ins.NumField())
	for i := 0; i < ins.NumField(); i++ {
		fields[typ.Field(i).Name] = ins.Field(i)
	}
	return fields
}

func BenchmarkReflectCacheNameSet(b *testing.B) {
	ref := ReflectCache(Config{})
	for i := 0; i < b.N; i++ {
		ref["Name"].SetString("name")
		ref["IP"].SetString("ip")
		ref["URL"].SetString("url")
		ref["Timeout"].SetString("timeout")
	}
}

func BenchmarkReferNameSet(b *testing.B) {
	fields := refer.Fields(&Config{})
	for i := 0; i < b.N; i++ {
		fields["Name"].SetString("name")
		fields["IP"].SetString("ip")
		fields["URL"].SetString("url")
		fields["Timeout"].SetString("timeout")
	}
}

func BenchmarkCall(b *testing.B) {
	c := Config{}
	for i := 0; i < b.N; i++ {
		c.Add(i, i)
	}
}

func BenchmarkReflectMethodIndexCall(b *testing.B) {
	value := reflect.ValueOf(Config{})

	for i := 0; i < b.N; i++ {
		value.Method(0).Call([]reflect.Value{reflect.ValueOf(i), reflect.ValueOf(i)})
	}
}

func BenchmarkReflectMethodNameCall(b *testing.B) {
	value := reflect.ValueOf(Config{})

	for i := 0; i < b.N; i++ {
		value.MethodByName("Add").Call([]reflect.Value{reflect.ValueOf(i), reflect.ValueOf(i)})
	}
}

func BenchmarkReferCall(b *testing.B) {
	methods := refer.Methods(Config{})

	for i := 0; i < b.N; i++ {
		methods["Add"].Call(refer.Args(i, i))
	}
}

func BenchmarkReferCopy(b *testing.B) {
	type Base struct {
		Name string
		Age  int
		Haha func() string
	}

	type Target struct {
		Name string
		Age  int
	}

	for i := 0; i < b.N; i++ {
		var base Base
		refer.Copy(&base, Target{Name: "dog", Age: 100})
	}
}

func BenchmarkReferCopyMap(b *testing.B) {
	type Base struct {
		Name string
		Age  int
		Haha func() string
	}

	type Target struct {
		Name string
		Age  int
	}

	for i := 0; i < b.N; i++ {
		var base Base
		if b, err := json.Marshal(Target{Name: "dog", Age: 100}); err == nil {
			_ = json.Unmarshal(b, &base)
		}
	}
}
