package loan

import (
	"errors"
	"loan-service/internal/constant"
	"loan-service/pkg/helper"
	"loan-service/pkg/model/request"
	"loan-service/pkg/model/response"
	"loan-service/pkg/responsehelper"
)

func (ls *Loan) GetList(status string, pagination *request.Pagination) (list []response.GetLoan, pageRes *response.Pagination) {
	st := constant.LoanStatus(status)
	loans := ls.loanRepo.GetList(st, pagination.Page, pagination.Size)
	total := ls.loanRepo.Count(st)

	list = make([]response.GetLoan, 0)

	for _, loan := range loans {
		list = append(list, response.GetLoan{
			ID:              loan.ID,
			BorrowerID:      loan.BorrowerID,
			Status:          loan.Status,
			PrincipalAmount: loan.PrincipalAmount,
			InvestedAmount:  loan.InvestedAmount,
			Rate:            loan.Rate,
			ROI:             loan.ROI,
			CreatedAt:       helper.FormatDate(loan.CreatedAt),
			CreatedBy:       loan.CreatedBy,
		})
	}

	return list, responsehelper.ToPagination(pagination, total)
}

func (ls *Loan) GetDetail(id string) (detail *response.GetLoanDetail, err error) {
	loanDetail := ls.loanRepo.GetDetail(id)

	if loanDetail == nil {
		return nil, errors.New("loan not exist")
	}

	investorsRes := make([]response.LoanInvestorDetail, 0)

	for _, investor := range loanDetail.LoanInvestment {
		investorsRes = append(investorsRes, response.LoanInvestorDetail{
			InvestorID: investor.InvestorID,
			Amount:     investor.Amount,
		})
	}

	return &response.GetLoanDetail{
		ID:              loanDetail.ID,
		BorrowerID:      loanDetail.BorrowerID,
		Status:          loanDetail.Status,
		PrincipalAmount: loanDetail.PrincipalAmount,
		InvestedAmount:  loanDetail.InvestedAmount,
		Rate:            loanDetail.Rate,
		ROI:             loanDetail.ROI,
		CreatedAt:       helper.FormatDate(loanDetail.CreatedAt),
		CreatedBy:       loanDetail.CreatedBy,
		Approval: &response.LoanApprovalDetail{
			FieldOfficerID: loanDetail.LoanApproval.FieldOfficerID,
			ProofOfPicture: loanDetail.LoanApproval.ProofOfPicture,
			Date:           helper.FormatDate(loanDetail.LoanApproval.CreatedAt),
		},
		Investors: investorsRes,
		Disbursement: &response.LoanDisbursementDetail{
			FieldOfficerID:          loanDetail.LoanDisbursement.FieldOfficerID,
			BorrowerAgreementLetter: loanDetail.LoanDisbursement.BorrowerAgreementLetter,
			Date:                    helper.FormatDate(loanDetail.LoanDisbursement.CreatedAt),
		},
	}, nil
}
