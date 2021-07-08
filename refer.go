package refer

import (
	"reflect"
)

type Refer struct {
	Fields    map[string]reflect.Value
	Methods   map[string]reflect.Value
	NumField  int
	NumMethod int
	Type      reflect.Type
	Value     *reflect.Value
}

func (r *Refer) Call(method string, args ...interface{}) ([]reflect.Value, error) {
	if r.Methods == nil {
		return nil, errCallEmptyMethods
	}

	if _, ok := r.Methods[method]; !ok {
		return nil, errCallNotHaveMethod{tip: method}
	}

	length := len(args)

	vals := make([]reflect.Value, length)

	for i := 0; i < length; i++ {
		vals[i] = reflect.ValueOf(args[i])
	}

	return r.Methods[method].Call(vals), nil
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

func New(target interface{}) *Refer {
	typ := reflect.TypeOf(target)
	ins := reflect.New(typ).Elem()
	numField := ins.NumField()
	numMethod := ins.NumMethod()
	r := Refer{
		Type:      typ,
		Value:     &ins,
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
	return &r
}
