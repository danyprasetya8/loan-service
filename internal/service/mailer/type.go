package mailer

import "bytes"

type Attachment struct {
	Name     string
	Content  *bytes.Buffer
	MimeType string
}

type Request struct {
	To         []string
	Subject    string
	Text       string
	Attachment *Attachment
}
