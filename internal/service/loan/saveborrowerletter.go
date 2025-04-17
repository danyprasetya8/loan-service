package loan

import (
	"errors"
	"loan-service/internal/constant"
	"loan-service/pkg/model/response"
	"mime/multipart"
)

func (ls *Loan) SaveBorrowerAgreementLetter(pdf *multipart.FileHeader, loanID, requestedBy string) (res response.UploadBorrowerLetter, err error) {
	loan := ls.loanRepo.Get(loanID)

	if loan == nil {
		return res, errors.New("loan not exist")
	}

	if loan.Status != constant.Invested {
		return res, errors.New("loan status must be invested")
	}

	pathPrefix := "borrowerAgreementLetter/" + loanID

	fileModel, err := ls.fileService.Save(pdf, constant.BorrowerAgreementLetter, pathPrefix, requestedBy)

	if err != nil {
		return
	}

	return response.UploadBorrowerLetter{
		ID:       fileModel.ID,
		Name:     fileModel.OriginalName,
		MimeType: fileModel.MimeType,
	}, nil
}
