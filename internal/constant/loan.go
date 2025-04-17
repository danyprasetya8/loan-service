package constant

type LoanStatus string

var (
	Proposed  LoanStatus = "proposed"
	Approved  LoanStatus = "approved"
	Invested  LoanStatus = "invested"
	Disbursed LoanStatus = "disbursed"
)
