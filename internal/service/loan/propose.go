package loan

import (
	"errors"
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/pkg/model/request"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

func (ls *Loan) Propose(req request.ProposeLoan, requestedBy string) (id string, err error) {
	borrower := ls.borrowerRepo.GetByID(req.BorrowerID)

	if borrower == nil {
		return "", errors.New("borrower not exist")
	}

	newLoan := &entity.Loan{
		ID:              uuid.New().String(),
		BorrowerID:      borrower.ID,
		Status:          constant.Proposed,
		PrincipalAmount: req.PrincipalAmount,
		InvestedAmount:  0,
		Rate:            req.Rate,
		ROI:             req.ROI,
		Audit: entity.Audit{
			CreatedBy: requestedBy,
			UpdatedBy: requestedBy,
		},
	}
	err = ls.loanRepo.Create(newLoan)

	if err != nil {
		log.Errorf("Error creating loan: %s", err.Error())
	}

	return newLoan.ID, err
}
