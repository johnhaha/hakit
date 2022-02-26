package hadata

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

//transfer version to int
func VersionToInt(v string, blockLength int) (int, error) {
	sp := strings.Split(VersionCleaner(v), ".")
	for i, s := range sp {
		for len(s) < blockLength {
			s = "0" + s
		}
		sp[i] = s
	}
	combine := NewStringBinder().BindString("1").BindString(sp...)
	res, err := GetIntFromString(combine.Value())
	return res, err
}

//rm v/V from head
func VersionCleaner(v string) string {
	s := strings.TrimLeft(v, "vV")
	return s
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

func (binder *StringBinder) BindWithSpace(str ...string) *StringBinder {
	for i, d := range str {
		binder.Builder.WriteString(d)
		if i < len(str)-1 {
			binder.Builder.WriteString(" ")
		}
	}
	return binder
}

func (binder *StringBinder) BindWithNewLine(str ...string) *StringBinder {
	for _, d := range str {
		binder.Builder.WriteString(d)
		binder.Builder.WriteString("\n")
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

//if v1 is empty, then v2
func EmptyThen(v1 string, v2 string) string {
	if v1 == "" {
		return v2
	}
	return v1
}

//generate random fixed length string
var upperLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var letters = []rune("abcdefghijklmnopqrstuvwxyz")
var digitalLetters = []rune("1234567890")

func GenerateRandomString(n int) string {
	var pool []rune
	pool = append(pool, upperLetters...)
	pool = append(pool, letters...)
	pool = append(pool, digitalLetters...)
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = pool[rand.Intn(len(pool))]
	}
	return string(b)
}

func GenerateRandomStringFromGivenRune(r []rune, n int) string {
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = r[rand.Intn(len(r))]
	}
	return string(b)
}
