package hadata

import "errors"

func Unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//check is data is in list
func IsInStringSlice(list []string, data string) bool {
	for _, l := range list {
		if l == data {
			return true
		}
	}
	return false
}

//swap slice
func SwapStringSlice(ds *[]string) error {
	if len(*ds) != 2 {
		return errors.New("not support")
	}
	dTemp := (*ds)[0]
	(*ds)[0] = (*ds)[1]
	(*ds)[1] = dTemp
	return nil
}
