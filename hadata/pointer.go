package hadata

import (
	"errors"
	"reflect"
)

func GetPointerData(data interface{}) (interface{}, error) {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		return v.Elem().Interface(), nil
	}
	return nil, errors.New("not a pointer")
}

func ClearPointer(data interface{}) interface{} {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		v, _ := GetPointerData(data)
		return v
	}
	return data
}
