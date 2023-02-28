package validator

import "strings"

func IsStringEmpty(str string) bool {
	return strings.Trim(str, " \t\n") == ""
}
