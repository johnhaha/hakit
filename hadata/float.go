package hadata

import (
	"fmt"
	"strconv"
)

func GetFloatFromString(d string) (float64, error) {
	f, err := strconv.ParseFloat(d, 64)
	return f, err
}

func GetStringFromFloat(f float64) string {
	return fmt.Sprintf("%f", f)
}
