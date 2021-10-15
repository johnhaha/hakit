package hadata

import (
	"errors"
	"strconv"
	"strings"
)

/*upgrade version
place start from 1
can be used in version like 1.0.1 or v1.0.1
*/
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
		op = op + "." + v
	}
	return op, nil
}
