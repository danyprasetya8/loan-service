package loan

import (
	"loan-service/internal/repository/borrower"
	"loan-service/internal/repository/loan"
	"loan-service/internal/repository/loanapproval"
	"loan-service/internal/repository/loandisbursement"
	"loan-service/internal/repository/loaninvestment"
	"loan-service/internal/repository/user"
	"loan-service/internal/service/file"
	"loan-service/pkg/model/request"
	"loan-service/pkg/model/response"
	"mime/multipart"
)

type ILoanService interface {
	Propose(req request.ProposeLoan, requestedBy string) (id string, err error)
	Approve(req request.ApproveLoan, requestedBy string) (success bool, err error)
	Invest(req request.InvestLoan, requestedBy string) (id string, err error)
	Disburse(req request.DisburseLoan, requestedBy string) (success bool, err error)
	SaveProofOfPicture(image *multipart.FileHeader, loanID, requestedBy string) (res response.UploadLoanProofOfPicture, err error)
	SaveBorrowerAgreementLetter(pdf *multipart.FileHeader, loanID, requestedBy string) (res response.UploadBorrowerLetter, err error)
}

type Dependency struct {
	FileService          file.IFileService
	UserRepo             user.IUserRepository
	BorrowerRepo         borrower.IBorrowerRepository
	LoanRepo             loan.ILoanRepository
	LoanApprovalRepo     loanapproval.ILoanApprovalRepository
	LoanInvestmentRepo   loaninvestment.ILoanInvestmentRepository
	LoanDisbursementRepo loandisbursement.ILoanDisbursementRepository
}

type Loan struct {
	fileService          file.IFileService
	userRepo             user.IUserRepository
	borrowerRepo         borrower.IBorrowerRepository
	loanRepo             loan.ILoanRepository
	loanApprovalRepo     loanapproval.ILoanApprovalRepository
	loanInvestmentRepo   loaninvestment.ILoanInvestmentRepository
	loanDisbursementRepo loandisbursement.ILoanDisbursementRepository
}

func New(deps *Dependency) ILoanService {
	return &Loan{
		fileService:          deps.FileService,
		userRepo:             deps.UserRepo,
		borrowerRepo:         deps.BorrowerRepo,
		loanRepo:             deps.LoanRepo,
		loanApprovalRepo:     deps.LoanApprovalRepo,
		loanInvestmentRepo:   deps.LoanInvestmentRepo,
		loanDisbursementRepo: deps.LoanDisbursementRepo,
	}
}
