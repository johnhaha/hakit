package hadata

import (
	"fmt"
	"math"
	"strconv"
)

type FloatData interface {
	float32 | float64
}

func ToFixed(data float64, place uint8) float64 {
	pw := math.Pow(10, float64(place))
	return math.Round(data*pw) / pw
}

func GetFloatFromString(d string) (float64, error) {
	f, err := strconv.ParseFloat(d, 64)
	return f, err
}

func GetStringFromFloat[T FloatData](f T) string {
	return fmt.Sprintf("%f", f)
}

/*
	data return

-1: f1<f2
0: f1=f2
1:f1>f2
*/
func FloatCompare(f1 float64, f2 float64, precision int) int {
	m := math.Pow10(precision)
	x := math.Round(f1*m) - math.Round(f2*m)
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}
