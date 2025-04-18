package response

import "loan-service/internal/constant"

type GetLoan struct {
	ID              string              `json:"id"`
	BorrowerID      string              `json:"borrowerId"`
	Status          constant.LoanStatus `json:"status"`
	PrincipalAmount int64               `json:"principalAmount"`
	InvestedAmount  int64               `json:"investedAmount"`
	Rate            float64             `json:"rate"`
	ROI             float64             `json:"roi"`
	CreatedAt       string              `json:"createdAt"`
	CreatedBy       string              `json:"createdBy"`
}

type LoanApprovalDetail struct {
	FieldOfficerID string `json:"fieldOfficerId"`
	ProofOfPicture string `json:"proofOfPicture"`
	Date           string `json:"date"`
}

type LoanInvestorDetail struct {
	InvestorID string `json:"id"`
	Amount     int64  `json:"amount"`
}

type LoanDisbursementDetail struct {
	FieldOfficerID          string `json:"fieldOfficerId"`
	BorrowerAgreementLetter string `json:"borrowerAgreementLetter"`
	Date                    string `json:"date"`
}

type GetLoanDetail struct {
	ID              string                  `json:"id"`
	BorrowerID      string                  `json:"borrowerId"`
	Status          constant.LoanStatus     `json:"status"`
	PrincipalAmount int64                   `json:"principalAmount"`
	InvestedAmount  int64                   `json:"investedAmount"`
	Rate            float64                 `json:"rate"`
	ROI             float64                 `json:"roi"`
	CreatedAt       string                  `json:"createdAt"`
	CreatedBy       string                  `json:"createdBy"`
	Approval        *LoanApprovalDetail     `json:"approval,omitempty"`
	Investors       []LoanInvestorDetail    `json:"investors"`
	Disbursement    *LoanDisbursementDetail `json:"disbursement,omitempty"`
}

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
