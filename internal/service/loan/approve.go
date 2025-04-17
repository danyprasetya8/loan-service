package loan

import (
	"errors"
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/pkg/model/request"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

func (ls *Loan) Approve(req request.ApproveLoan, requestedBy string) (success bool, err error) {
	loan := ls.loanRepo.Get(req.LoanID)

	if loan == nil {
		return false, errors.New("loan not exist")
	}

	if loan.Status != constant.Proposed {
		return false, errors.New("loan status must be proposed")
	}

	fieldOfficer := ls.userRepo.Get(req.FieldOfficerID)

	if fieldOfficer == nil {
		return false, errors.New("field officer not exist")
	}

	if fieldOfficer.Role != constant.FieldOfficer {
		return false, errors.New("invalid user role")
	}

	fileModel := ls.fileService.Find(req.ProofOfPicture)

	if fileModel == nil {
		return false, errors.New("proof of picture not exist")
	}

	if fileExist := ls.fileService.IsExist(fileModel.Path); !fileExist {
		return false, errors.New("proof of picture not exist")
	}

	loanApproval := &entity.LoanApproval{
		ID:             uuid.New().String(),
		LoanID:         loan.ID,
		FieldOfficerID: fieldOfficer.Email,
		ProofOfPicture: fileModel.ID,
		Audit: entity.Audit{
			CreatedBy: requestedBy,
			UpdatedBy: requestedBy,
		},
	}
	if err = ls.loanApprovalRepo.Create(loanApproval); err != nil {
		log.Errorf("Error creating loan approval: %s", err.Error())
		return
	}

	loan.Status = constant.Approved
	ls.loanRepo.Save(loan)

	return true, nil
}
