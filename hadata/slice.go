package hadata

import "errors"

func Unique[T comparable](strSlice []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// check is data is in list
func IsInSlice[T comparable](list []T, data T) bool {
	for _, l := range list {
		if l == data {
			return true
		}
	}
	return false
}

//get inter of two slice

func InterSlice[T comparable](s1 []T, s2 []T) []T {
	var op []T
	for _, s := range s1 {
		if IsInSlice(s2, s) {
			op = append(op, s)
		}
	}
	return op
}

// swap slice
func SwapSlice[T any](ds *[]T) error {
	if len(*ds) != 2 {
		return errors.New("not support")
	}
	dTemp := (*ds)[0]
	(*ds)[0] = (*ds)[1]
	(*ds)[1] = dTemp
	return nil
}

func Map[T any, Q any](data []T, trans func(T) Q) []Q {
	l := len(data)
	ot := make([]Q, l)
	for i := 0; i < l; i++ {
		ot[i] = trans(data[i])
	}
	return ot
}

func Maps[T any, Q any, R any](data []T, trans func(T) (Q, R)) ([]Q, []R) {
	l := len(data)
	ot1 := make([]Q, l)
	ot2 := make([]R, l)
	for i := 0; i < l; i++ {
		ot1[i], ot2[i] = trans(data[i])
	}
	return ot1, ot2
}

func Transform[T any, Q any](data []T, trans func(T) []Q) []Q {
	var ot []Q
	for _, d := range data {
		ot = append(ot, trans(d)...)
	}
	return ot
}

func Combine[T any, Q any](data []T, target Q, combine func(T, Q) T) []T {
	for i, d := range data {
		data[i] = combine(d, target)
	}
	return data
}

func RemElement[T any](data []T, index int) []T {
	return append(data[:index], data[index+1:]...)
}

// remove specific element by replace it with the last element, it's faster, but order changed
func RemElementX[T any](data []T, index int) []T {
	l := len(data)
	data[index] = data[l-1]
	return data[:l-1]
}

func RemWhere[T any](data []T, where func(T) bool) []T {
	if l := len(data); l > 0 {
		var ot []T
		for i := 0; i < l; i++ {
			check := data[i]
			if where(check) {
				continue
			}
			ot = append(ot, check)
		}
		return ot
	}
	return nil
}

// remove element by replace it with the last elements, it's faster, but order will be changed
//
// use this function if you don't care about the original order
func RemWhereX[T any](data []T, where func(T) bool) []T {
	if l := len(data); l > 0 {
		var k int
		for i := l - 1; i >= 0; i-- {
			check := data[i]
			if where(check) {
				k++
				data[i] = data[l-k]
			}
		}
		return data[:l-k]
	}
	return nil
}

func Fold[V any, T any](data []T, initValue V, fold func(v V, d T) V) V {
	value := initValue
	if l := len(data); l > 0 {
		for i := 0; i < l; i++ {
			value = fold(value, data[i])
		}
	}
	return value
}
