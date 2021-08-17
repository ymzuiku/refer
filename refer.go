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
	vals := make([]reflect.Value, len(args))

	for i, v := range args {
		vals[i] = reflect.ValueOf(v)
	}

	return vals
}

func Call(target interface{}, key string, args ...interface{}) []reflect.Value {
	method := reflect.ValueOf(target).MethodByName(key)
	if !method.IsNil() {
		return method.Call(Args(args))
	}

	return []reflect.Value{}
}

func Set(target interface{}, key string, value interface{}) {
	field := reflect.ValueOf(target).FieldByName(key)
	if !field.IsNil() {
		field.Set(reflect.ValueOf(value))
	}
}

func Get(target interface{}, key string) reflect.Value {
	return reflect.ValueOf(target).FieldByName(key)
}
