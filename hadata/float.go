package hadata

import "strconv"

func GetFloatFromString(d string) (float64, error) {
	f, err := strconv.ParseFloat(d, 64)
	return f, err
}
