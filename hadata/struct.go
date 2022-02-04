package hadata

import (
	"errors"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

func StructToMap(item any) map[string]any {

	res := map[string]any{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = StructToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}

func StructToStringMap(item any) map[string]string {
	res := map[string]string{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).String()
		res[tag] = field
	}
	return res
}

func MapToStruct(input any, output any) error {
	err := mapstructure.Decode(input, output)
	return err
}

func GetStructName(data any) string {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}

func GetStructNameInLowerCase(data any) string {
	t := reflect.TypeOf(data)
	var name string
	if t.Kind() == reflect.Ptr {
		name = t.Elem().Name()
	} else {
		name = t.Name()
	}
	return strings.ToLower(name)
}

//return struct tag data, in json key, empty field will be dropped unless specified in including field
func ReadStructTagData(data any, tag string, includingField ...string) (map[string]any, error) {
	d := ClearPointer(data)

	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	mp := make(map[string]any)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if _, ok := f.Tag.Lookup(tag); ok {
			k, _ := getJsonFieldName(f)
			fv := v.Field(i)
			if fv.IsZero() && !IsInSlice(includingField, k) {
				continue
			}
			mp[k] = v.Field(i).Interface()
		}
	}
	return mp, nil
}

func LookUpFirstTagMark(data any, tag string, mark string) (name string, value any, err error) {
	d := ClearPointer(data)

	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		res := getFiledTagSlice(f, tag)
		if res != nil && IsInSlice(res, mark) {
			name, _ := getJsonFieldName(f)
			return name, v.Field(i).Interface(), nil
		}
	}
	return "", nil, errors.New("not found")
}
