package hadata

import (
	"errors"
	"reflect"
)

func GetPointerData(data any) (any, error) {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		return v.Elem().Interface(), nil
	}
	return nil, errors.New("not a pointer")
}

func ClearPointer(data any) any {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		v, _ := GetPointerData(data)
		return v
	}
	return data
}
