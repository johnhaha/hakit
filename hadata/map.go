package hadata

import "encoding/json"

// transfer 'any' value to string or json if not string
func GetStringMap(data map[string]any) (map[string]string, error) {
	ot := make(map[string]string)
	for k, v := range data {
		switch d := v.(type) {
		case string:
			ot[k] = d
		default:
			x, err := json.Marshal(d)
			if err != nil {
				return nil, err
			}
			ot[k] = string(x)
		}
	}
	return ot, nil
}

func MapFromSlice[T any](data []T, getKey func(T) string) map[string]T {
	ot := make(map[string]T)
	for _, d := range data {
		ot[getKey(d)] = d
	}
	return ot
}

func GetMapKey(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
