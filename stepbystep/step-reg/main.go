///www.cnblogs.com/williamjie/p/9686311.html

package main

import (
	"fmt"
	"regexp"
)

func IsEmail(email string) bool {
	if email != "" {
		if isOk, _ := regexp.MatchString("^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$", email); isOk {
			return true
		}
	}

	return false
}

func IsPhone(phoneStr string) bool {
	if phoneStr != "" {
		if isOk, _ := regexp.MatchString(`^\([\d]{3}\) [\d]{3}-[\d]{4}$`, phoneStr); isOk {
			return isOk
		}
	}

	return false
}


func IsStr(str string) bool {
	if str != "" {
		if isOk, _ := regexp.MatchString(`^[1-9]\d*s$`, str); isOk {

			return isOk
		}
	}

	return false
}


func FindNum() {
	text := `123s4s`
	// 查找连续的小写字母
	reg := regexp.MustCompile(`^[1-9]\d*s`)
	fmt.Printf("%s\n", reg.FindString(text))
}

func main() {
	FindNum()
}