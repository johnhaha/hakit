package hadata

import (
	"math/rand"
	"time"
)

//return random int between min and max, not include max
func GetRandomNumber(min int, max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n := r.Intn(max - min)
	return n + min
}

//return random int between min and max, not include max
func GetManyRandomNumber(min int, max int, count int) []int {
	ot := make([]int, count)
	for i := 0; i < count; i++ {
		ot[i] = GetRandomNumber(min, max)
	}
	return ot
}
