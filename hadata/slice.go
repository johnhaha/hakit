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

//check is data is in list
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

//swap slice
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
	ot := make([]Q, len(data))
	for i, d := range data {
		ot[i] = trans(d)
	}
	return ot
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
