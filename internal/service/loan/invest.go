package loan

import (
	"errors"
	"fmt"
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/pkg/model/request"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

func (ls *Loan) Invest(req request.InvestLoan, requestedBy string) (success bool, err error) {
	loan := ls.loanRepo.Get(req.LoanID)

	if loan == nil {
		return false, errors.New("loan not exist")
	}

	if loan.Status != constant.Approved {
		return false, errors.New("loan status must be approved")
	}

	newAmount := loan.InvestedAmount + req.Amount

	if newAmount > loan.PrincipalAmount {
		remainingNeeded := loan.PrincipalAmount - loan.InvestedAmount
		return false, fmt.Errorf("remaining amount needed: %d", remainingNeeded)
	}

	existLoanInvestment := ls.loanInvestmentRepo.GetByLoanAndInvestor(loan.ID, requestedBy)

	if existLoanInvestment != nil {
		existLoanInvestment.Amount += req.Amount
		if err = ls.loanInvestmentRepo.Save(existLoanInvestment); err != nil {
			log.Errorf("Error saving loan investment: %s", err.Error())
			return
		}
	} else {
		newInvestment := &entity.LoanInvestment{
			ID:         uuid.New().String(),
			LoanID:     loan.ID,
			InvestorID: requestedBy,
			Amount:     req.Amount,
			Audit: entity.Audit{
				CreatedBy: requestedBy,
				UpdatedBy: requestedBy,
			},
		}
		if err = ls.loanInvestmentRepo.Create(newInvestment); err != nil {
			log.Errorf("Error creating loan investment: %s", err.Error())
			return
		}
	}

	loan.InvestedAmount += req.Amount
	if loan.InvestedAmount == loan.PrincipalAmount {
		loan.Status = constant.Invested
	}

	if err = ls.loanRepo.Save(loan); err != nil {
		log.Errorf("Error saving loan: %s", err.Error())
		return
	}

	return true, nil
}
