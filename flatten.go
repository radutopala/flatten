package flatten

import (
	"reflect"
)

func It(v ...interface{}) []interface{} {
	return flatten(nil, reflect.ValueOf(v))
}

func flatten(args []interface{}, v reflect.Value) []interface{} {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			args = flatten(args, v.Index(i))
		}
	} else {
		args = append(args, v.Interface())
	}
	return args
}
