package hadata

import "encoding/json"

func GetDataFromMap[T any](m interface{}) (T, error) {
	var data T
	err := MapToStruct(m, &data)
	return data, err
}

func GetDataFromJson[T any](m string) (T, error) {
	var data T
	err := json.Unmarshal([]byte(m), &data)
	return data, err
}
