package hadata

import (
	"reflect"
	"strings"
)

func GetJsonFieldName(f reflect.StructField) (name string, omitempty bool) {
	if n, ok := f.Tag.Lookup("json"); ok {
		ns := strings.Split(n, ",")
		if len(ns) == 1 {
			return n, false
		}
		return ns[0], true
	}
	return f.Name, false
}

func GetFiledTagSlice(f reflect.StructField, tag string) []string {
	if n, ok := f.Tag.Lookup(tag); ok {
		s := strings.Split(n, ",")
		return s
	}
	return nil
}
