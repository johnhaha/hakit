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
