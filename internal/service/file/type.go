package file

import "loan-service/internal/constant"

type Model struct {
	ID           string
	OriginalName string
	Path         string
	MimeType     string
	Type         constant.FileType
}
