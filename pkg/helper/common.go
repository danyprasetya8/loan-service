package helper

import (
	"loan-service/internal/constant"
	"mime/multipart"
	"net/http"
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

func GetMimeType(fileHeader *multipart.FileHeader) (m string, err error) {
	file, err := fileHeader.Open()

	if err != nil {
		return
	}

	defer file.Close()

	buffer := make([]byte, 512)
	if _, err = file.Read(buffer); err != nil {
		return
	}

	mimeType := http.DetectContentType(buffer)

	return mimeType, nil
}

func SplitLast(str, sep string) (s1 string, s2 string) {
	idx := strings.LastIndex(str, sep)
	if idx == -1 {
		return str, ""
	}
	return str[:idx], str[idx+1:]
}
