package helpers

import (
	"regexp"
)

func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^0[0-9]{9}$`)
	return re.MatchString(phone)
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValidateLength(value string, min, max int) bool {
	l := len(value)
	return l >= min && l <= max
}