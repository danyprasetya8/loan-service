package helper

import "strings"

func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}
