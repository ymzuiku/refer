package refer

import (
	"reflect"
)

type Refer struct {
	kind      reflect.Kind
	Fields    map[string]reflect.Value
	Methods   map[string]reflect.Value
	NumField  int
	NumMethod int
	Type      reflect.Type
	Value     reflect.Value
}

func (r *Refer) Call(method string, args ...interface{}) []reflect.Value {
	if r.Methods == nil {
		return nil
	}

	if _, ok := r.Methods[method]; !ok {
		return nil
	}

	length := len(args)

	vals := make([]reflect.Value, length)

	for i := 0; i < length; i++ {
		vals[i] = reflect.ValueOf(args[i])
	}

	return r.Methods[method].Call(vals)
}

func (r *Refer) F(key string) reflect.Value {
	return r.Fields[key]
}

func (r *Refer) GetFieldNames() []string {
	list := make([]string, 0, r.NumField)
	for k := range r.Fields {
		list = append(list, k)
	}
	return list
}

func (r *Refer) GetMethodNames() []string {
	list := make([]string, 0, r.NumMethod)
	for k := range r.Methods {
		list = append(list, k)
	}
	return list
}

func (r *Refer) Interface() interface{} {
	return r.Value.Interface()
}

func New(target interface{}) Refer {
	typ := reflect.TypeOf(target)
	kind := typ.Kind()
	if kind != reflect.Struct {
		typ = typ.Elem()
	}
	ins := reflect.ValueOf(target)
	if kind != reflect.Struct {
		ins = ins.Elem()
	}
	numField := ins.NumField()
	numMethod := ins.NumMethod()
	r := Refer{
		kind:      kind,
		Type:      typ,
		Value:     ins,
		NumField:  numField,
		NumMethod: numMethod,
	}
	if numField > 0 {
		r.Fields = make(map[string]reflect.Value, numField)
		for i := 0; i < numField; i++ {
			r.Fields[typ.Field(i).Name] = ins.Field(i)
		}
	}

	if numMethod > 0 {
		r.Methods = make(map[string]reflect.Value, numMethod)
		for i := 0; i < numMethod; i++ {
			r.Methods[typ.Method(i).Name] = ins.Method(i)
		}
	}
	return r
}

// func Copy(base interface{}, target interface{}) {
// 	t := New(target)
// 	b := New(base)

// 	if t.NumField > 0 && b.NumField > 0 {
// 		for k, v := range t.Fields {
// 			if _, ok := b.Fields[k]; ok {
// 				b.Fields[k].Set(v)
// 			}
// 		}
// 	}

// 	if t.NumMethod > 0 && b.NumMethod > 0 {
// 		for k, v := range t.Methods {
// 			if _, ok := b.Methods[k]; ok {
// 				b.Methods[k].Set(v)
// 			}
// 		}
// 	}
// }
