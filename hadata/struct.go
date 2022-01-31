package hadata

import (
	"errors"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

func StructToMap(item interface{}) map[string]interface{} {

	res := map[string]interface{}{}
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

func StructToStringMap(item interface{}) map[string]string {
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

func MapToStruct(input interface{}, output interface{}) error {
	err := mapstructure.Decode(input, output)
	return err
}

func GetStructName(data interface{}) string {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}

func GetStructNameInLowerCase(data interface{}) string {
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
func ReadStructTagData(data interface{}, tag string, includingField ...string) (map[string]interface{}, error) {
	d := ClearPointer(data)

	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	mp := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if _, ok := f.Tag.Lookup(tag); ok {
			k, _ := getJsonFieldName(f)
			fv := v.Field(i)
			if fv.IsZero() && !IsInStringSlice(includingField, k) {
				continue
			}
			mp[k] = v.Field(i).Interface()
		}
	}
	return mp, nil
}

func LookUpFirstTagMark(data interface{}, tag string, mark string) (name string, value interface{}, err error) {
	d := ClearPointer(data)

	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		res := getFiledTagSlice(f, tag)
		if res != nil && IsInStringSlice(res, mark) {
			name, _ := getJsonFieldName(f)
			return name, v.Field(i).Interface(), nil
		}
	}
	return "", nil, errors.New("not found")
}
