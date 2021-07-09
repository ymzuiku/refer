package unit

import (
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

func ReflectHelp(target interface{}) map[string]reflect.Value {
	typ := reflect.TypeOf(target)
	ins := reflect.New(typ).Elem()
	fields := make(map[string]reflect.Value, ins.NumField())
	for i := 0; i < ins.NumField(); i++ {
		fields[typ.Field(i).Name] = ins.Field(i)
	}
	return fields
}

func BenchmarkReflectHelpNameSet(b *testing.B) {
	ref := ReflectHelp(Config{})
	for i := 0; i < b.N; i++ {
		ref["Name"].SetString("name")
		ref["IP"].SetString("ip")
		ref["URL"].SetString("url")
		ref["Timeout"].SetString("timeout")
	}
}

func BenchmarkReflectReferNameSet(b *testing.B) {
	ref := refer.New(Config{})
	for i := 0; i < b.N; i++ {
		ref.Fields["Name"].SetString("name")
		ref.Fields["IP"].SetString("ip")
		ref.Fields["URL"].SetString("url")
		ref.Fields["Timeout"].SetString("timeout")
	}
}

func BenchmarkReflectReferNameUseGet(b *testing.B) {
	ref := refer.New(Config{})
	for i := 0; i < b.N; i++ {
		ref.F("Name").SetString("name")
		ref.F("IP").SetString("ip")
		ref.F("URL").SetString("url")
		ref.F("Timeout").SetString("timeout")
	}
}

func BenchmarkAdd(b *testing.B) {
	c := Config{}
	for i := 0; i < b.N; i++ {
		c.Add(i, i)
	}
}

func BenchmarkReflectAdd(b *testing.B) {
	ref := refer.New(Config{})

	for i := 0; i < b.N; i++ {
		ref.Call("Add", i, i)
	}
}
