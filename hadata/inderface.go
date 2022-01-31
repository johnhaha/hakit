package hadata

import (
	"errors"
	"reflect"
)

func GetStringFromInterface(x interface{}) string {
	switch v := x.(type) {
	case string:
		return v
	default:
		return ""
	}
}

func CheckInterfaceDataIsSlice(x interface{}) (bool, interface{}) {
	data := ClearPointer(x)
	return reflect.TypeOf(data).Kind() == reflect.Slice, data
}

func GetSliceFromInterface(x interface{}) ([]interface{}, error) {
	check, data := CheckInterfaceDataIsSlice(x)
	if !check {
		return nil, errors.New("not a slice")
	}
	s := reflect.ValueOf(data)
	ot := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ot[i] = s.Index(i).Interface()
	}
	return ot, nil
}
