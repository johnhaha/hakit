package haerr

func FailOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintOnError(err error) {
	if err != nil {
		println(err)
	}
}

func CheckError(err error) bool {
	if err != nil {
		println(err.Error())
	}
	return err != nil
}
