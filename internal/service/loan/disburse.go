package loan

import (
	"errors"
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/pkg/model/request"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

func (ls *Loan) Disburse(req request.DisburseLoan, requestedBy string) (success bool, err error) {
	loan := ls.loanRepo.Get(req.LoanID)

	if loan == nil {
		return false, errors.New("loan not exist")
	}

	if loan.Status != constant.Invested {
		return false, errors.New("loan status must be invested")
	}

	fieldOfficer := ls.userRepo.Get(req.FieldOfficerID)

	if fieldOfficer == nil {
		return false, errors.New("field officer not exist")
	}

	if fieldOfficer.Role != constant.FieldOfficer {
		return false, errors.New("invalid user role")
	}

	fileModel := ls.fileService.Find(req.BorrowerAgreementLetter)

	if fileModel == nil {
		return false, errors.New("agreement letter not exist")
	}

	if fileExist := ls.fileService.IsExist(fileModel.Path); !fileExist {
		return false, errors.New("agreement letter not exist")
	}

	loandisbursement := &entity.LoanDisbursement{
		ID:                      uuid.New().String(),
		LoanID:                  loan.ID,
		FieldOfficerID:          fieldOfficer.Email,
		BorrowerAgreementLetter: fileModel.ID,
		Audit: entity.Audit{
			CreatedBy: requestedBy,
			UpdatedBy: requestedBy,
		},
	}
	if err = ls.loanDisbursementRepo.Create(loandisbursement); err != nil {
		log.Errorf("Error creating loan disbursement: %s", err.Error())
		return
	}

	loan.Status = constant.Disbursed
	ls.loanRepo.Save(loan)

	return true, nil
}
