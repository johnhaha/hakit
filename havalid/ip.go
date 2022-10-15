package havalid

import (
	"errors"
	"net"

	"github.com/asaskevich/govalidator"
)

// check if ip and type, return 4 and 6 as valid type, return 0 as not ip
func IPCheck(ip string) int {
	is4 := govalidator.IsIPv4(ip)
	if is4 {
		return 4
	}
	is6 := govalidator.IsIPv6(ip)
	if is6 {
		return 6
	}
	return 0
}

func IsPrivateIP(ip string) (bool, error) {
	if IPCheck(ip) > 0 {
		if ipAddr := net.ParseIP(ip); ipAddr != nil {
			return ipAddr.IsPrivate(), nil
		}
	}
	return false, errors.New("not ip")
}
