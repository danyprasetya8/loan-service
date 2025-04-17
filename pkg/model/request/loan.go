package request

type ProposeLoan struct {
	BorrowerID      string  `json:"borrowerId"`
	PrincipalAmount int64   `json:"principalAmount"`
	Rate            float64 `json:"rate"`
	ROI             float64 `json:"roi"`
}

type ApproveLoan struct {
	LoanID         string `json:"-"`
	FieldOfficerID string `json:"fieldOfficerId"`
	ProofOfPicture string `json:"proofOfPicture"`
}

type InvestLoan struct {
	LoanID string `json:"-"`
	Amount int64  `json:"amount"`
}

type DisburseLoan struct {
	LoanID                  string `json:"-"`
	FieldOfficerID          string `json:"fieldOfficerId"`
	BorrowerAgreementLetter string `json:"borrowerAgreementLetter"`
}
