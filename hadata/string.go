package hadata

import (
	"errors"
	"strconv"
	"strings"
)

/* upgrade version
place start from 1
 can be used in version like 1.0.1 or v1.0 */
func UpgradeVersion(version string, place int) (string, error) {
	sp := strings.Split(version, ".")
	if len(sp) < place {
		return "", errors.New("bad version")
	}
	withHead := false
	if place == 1 {
		if sp[0][0:1] == "v" {
			withHead = true
			sp[0] = sp[0][1:]
		}
	}
	l, err := strconv.Atoi(sp[place-1])
	if err != nil {
		return "", err
	}
	sp[place-1] = strconv.Itoa(l + 1)
	op := ""
	if withHead {
		op = "v"
	}
	for i, v := range sp {
		if i == 0 {
			op = op + v
			continue
		}
		if i >= place {
			op = op + ".0"
			continue
		}
		op = op + "." + v
	}
	return op, nil
}

//get string from int
func GetStringFromInt(d int) string {
	return strconv.Itoa(d)
}

func GetIntFromString(s string) (int, error) {
	intVar, err := strconv.Atoi(s)
	return intVar, err
}

type StringBinder struct {
	Builder strings.Builder
}

func NewStringBinder() *StringBinder {
	return &StringBinder{}
}

func (binder *StringBinder) BindString(str ...string) *StringBinder {

	for _, d := range str {
		binder.Builder.WriteString(d)
	}
	return binder
}

func (binder *StringBinder) BindInt(num ...int) *StringBinder {
	for _, d := range num {
		binder.Builder.WriteString(strconv.Itoa(d))
	}
	return binder
}

func (binder *StringBinder) Value() string {
	return binder.Builder.String()
}
