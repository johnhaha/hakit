package havalid

import "github.com/asaskevich/govalidator"

//check if ip and type, return 4 and 6 as valid type, return 0 as not ip
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
