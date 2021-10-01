package haerr

func FailOnError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func PrintOnError(err error, msg string) {
	if err != nil {
		println(msg)
	}
}

func CheckError(err error) bool {
	if err != nil {
		println(err.Error())
	}
	return err != nil
}
