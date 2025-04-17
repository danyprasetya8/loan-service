package response

type UploadLoanProofOfPicture struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	MimeType string `json:"type"`
}

type UploadBorrowerLetter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	MimeType string `json:"type"`
}
