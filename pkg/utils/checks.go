package utils

import (
	"regexp"
	"strings"
)

func CheckValidity(login_type string, identifier string) (bool, string) {
	switch login_type {
	case "mobile":
		matched, _ := regexp.MatchString(`^\d{10}$`, identifier)
		if !matched {
			return matched, "Invalid mobile no. Must be 10 digits."
		}
	case "email":
		if !strings.Contains(identifier, "@") || !strings.Contains(identifier, ".") {
			return false, "Invalid email format."
		}
	default:
		return false, "Invalid login type"
	}
	return true, ""
}
