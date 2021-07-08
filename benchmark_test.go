package refer

import (
	"reflect"
	"testing"
)

type Config struct {
	Name    string `json:"server-name"`
	IP      string `json:"server-ip"`
	URL     string `json:"server-url"`
	Timeout string `json:"timeout"`
}

// func readConfig() *Config {
// 	config := Config{}
// 	typ := reflect.TypeOf(config)
// 	value := reflect.Indirect(reflect.ValueOf(&config))

// 	for i := 0; i < typ.NumField(); i++ {
// 		f := typ.Field(i)
// 		if v, ok := f.Tag.Lookup("json"); ok {
// 			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
// 			if env, exist := os.LookupEnv(key); exist {
// 				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
// 			}
// 		}
// 	}

// 	return &config
// }

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
	b.ResetTimer()
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
	b.ResetTimer()
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
	b.ResetTimer()
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
	ref := New(Config{})
	for i := 0; i < b.N; i++ {
		ref.Fields["Name"].SetString("name")
		ref.Fields["IP"].SetString("ip")
		ref.Fields["URL"].SetString("url")
		ref.Fields["Timeout"].SetString("timeout")
	}
}

func BenchmarkReflectReferNameUseGet(b *testing.B) {
	ref := New(Config{})
	for i := 0; i < b.N; i++ {
		ref.F("Name").SetString("name")
		ref.F("IP").SetString("ip")
		ref.F("URL").SetString("url")
		ref.F("Timeout").SetString("timeout")
	}
}
