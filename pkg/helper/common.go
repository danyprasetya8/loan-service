package helper

import (
	"loan-service/internal/constant"
	"strings"
	"time"
)

func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

func FormatDate(dt time.Time) string {
	loc, _ := time.LoadLocation(constant.AsiaJakarta)
	return dt.In(loc).Format(constant.DateLayout)
}

func IsValidEmail(str string) bool {
	return constant.EmailRegex.MatchString(str)
}
