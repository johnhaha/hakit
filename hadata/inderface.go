package hadata

import "reflect"

func GetStringFromInterface(x interface{}) string {
	switch v := x.(type) {
	case string:
		return v
	default:
		return ""
	}
}

func CheckInterfaceIsSlice(x interface{}) bool {
	return reflect.TypeOf(x).Kind() == reflect.Slice
}
