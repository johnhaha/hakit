package hadata

import (
	"fmt"
	"strconv"
)

type FloatData interface {
	float32 | float64
}

func GetFloatFromString(d string) (float64, error) {
	f, err := strconv.ParseFloat(d, 64)
	return f, err
}

func GetStringFromFloat[T FloatData](f T) string {
	return fmt.Sprintf("%f", f)
}
