package hadata

func GetStringFromInterface(x interface{}) string {
	switch v := x.(type) {
	case string:
		return v
	default:
		return ""
	}
}
