package refer

import (
	"reflect"
)

func MethodsList(target interface{}) []string {
	typ := reflect.TypeOf(target)
	if typ.Kind() != reflect.Struct {
		typ = typ.Elem()
	}
	num := typ.NumMethod()
	fields := []string{}

	for i := 0; i < num; i++ {
		fields = append(fields, typ.Method(i).Name)
	}

	return fields
}

func Methods(target interface{}) map[string]reflect.Value {
	typ := reflect.TypeOf(target)
	ins := reflect.ValueOf(target)
	if typ.Kind() != reflect.Struct {
		typ = typ.Elem()
		ins = ins.Elem()
	}
	num := typ.NumMethod()

	fields := make(map[string]reflect.Value, num)
	for i := 0; i < num; i++ {
		fields[typ.Method(i).Name] = ins.Method(i)
	}

	return fields
}

func Fields(target interface{}) map[string]reflect.Value {
	typ := reflect.TypeOf(target)
	ins := reflect.ValueOf(target)
	if typ.Kind() != reflect.Struct {
		typ = typ.Elem()
		ins = ins.Elem()
	}
	num := typ.NumField()

	out := make(map[string]reflect.Value, num)
	for i := 0; i < num; i++ {
		out[typ.Field(i).Name] = ins.Field(i)
	}

	return out
}

func FieldsList(target interface{}) []string {
	typ := reflect.TypeOf(target)
	if typ.Kind() != reflect.Struct {
		typ = typ.Elem()
	}
	num := typ.NumField()

	fields := []string{}
	for i := 0; i < num; i++ {
		fields = append(fields, typ.Field(i).Name)
	}
	return fields
}

func Copy(base interface{}, target interface{}) {
	b := Fields(base)
	if b != nil {
		t := Fields(target)
		for k, v := range t {
			if _, ok := b[k]; ok {
				b[k].Set(v)
			}
		}
	}
}

func Args(args ...interface{}) []reflect.Value {
	l := len(args)
	if l == 0 {
		return []reflect.Value{}
	}
	vals := make([]reflect.Value, l)

	for i, v := range args {
		vals[i] = reflect.ValueOf(v)
	}

	return vals
}

func Call(target interface{}, name string, args ...interface{}) []reflect.Value {
	return reflect.ValueOf(target).MethodByName(name).Call(Args(args...))
}

func Set(target interface{}, name string, value interface{}) {
	reflect.ValueOf(target).Elem().FieldByName(name).Set(reflect.ValueOf(value))
}

func Get(target interface{}, name string) reflect.Value {
	return reflect.ValueOf(target).FieldByName(name)
}
