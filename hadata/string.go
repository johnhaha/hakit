package hadata

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetFirstNotEmptyString(candidates ...string) (value string) {
	for _, v := range candidates {
		if v != "" {
			return v
		}
	}
	return
}

func TitleWord(word string) string {
	c := cases.Title(language.English, cases.NoLower)
	return c.String(word)
}

/*
	upgrade version

place start from 1

	can be used in version like 1.0.1 or v1.0
*/
func UpgradeVersion(version string, place int) (string, error) {
	head, body := GetVersionData(version)

	sp := strings.Split(body, ".")
	if len(sp) < place {
		return "", errors.New("bad version")
	}
	// head := ""

	// if place == 1 {
	// 	if h := sp[0][0:1]; h == "v" || h == "^" {
	// 		head = h
	// 		sp[0] = sp[0][1:]
	// 	}
	// }
	l, err := strconv.Atoi(sp[place-1])
	if err != nil {
		return "", err
	}
	sp[place-1] = strconv.Itoa(l + 1)

	op := head
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

func GetVersionData(version string) (head string, body string) {
	if version == "" {
		return
	}
	if h := version[0:1]; h == "v" || h == "^" {
		head = h
		body = version[1:]
		return
	}
	return "", version
}

// transfer version to int
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

// rm v/V from head
func VersionCleaner(v string) string {
	s := strings.TrimLeft(v, "vV")
	return s
}

// get string from int
func GetStringFromInt(d int) string {
	return strconv.Itoa(d)
}

func GetIntFromString(s string) (int, error) {
	intVar, err := strconv.Atoi(s)
	return intVar, err
}

func GetInt64FromString(s string) (int64, error) {
	intVar, err := strconv.ParseInt(s, 10, 64)
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

// if v1 is empty, then v2
func EmptyThen(v1 string, v2 string) string {
	if v1 == "" {
		return v2
	}
	return v1
}

// generate random fixed length string
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

func SplitStringOnUpperCase(str string) []string {
	var ot []string
	var e string
	for i, s := range str {
		if i == 0 {
			e = string(s)
			continue
		}
		if unicode.IsUpper(s) {
			ot = append(ot, e)
			e = string(s)
		} else {
			e += string(s)
		}
	}
	ot = append(ot, e)
	return ot
}

func CaseCamelToSnake(str string) string {
	res := SplitStringOnUpperCase(str)
	var ot string
	for i, s := range res {
		x := strings.ToLower(s)
		if i == 0 {
			ot = x
			continue
		}
		ot += "_" + x
	}
	return ot
}

func SnakeToCaseCamel(str string, isTitle bool) string {
	sp := strings.Split(str, "_")
	var ot string
	for i, s := range sp {
		if i == 0 && !isTitle {
			ot = s
			continue
		}
		ot += TitleWord(s)
	}
	return ot
}

func RuneLen(str string) int {
	return len([]rune(str))
}
