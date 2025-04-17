package loan

import (
	"errors"
	"loan-service/internal/constant"
	"loan-service/pkg/model/response"
	"mime/multipart"
	"path/filepath"
)

func (ls *Loan) SaveProofOfPicture(image *multipart.FileHeader, loanID, requestedBy string) (res response.UploadLoanProofOfPicture, err error) {
	loan := ls.loanRepo.Get(loanID)

	if loan == nil {
		return res, errors.New("loan not exist")
	}

	if loan.Status != constant.Proposed {
		return res, errors.New("loan status must be proposed")
	}

	pathPrefix := filepath.Join("proofOfPicture", loanID)

	fileModel, err := ls.fileService.Save(image, constant.ProofOfPicture, pathPrefix, requestedBy)

	if err != nil {
		return
	}

	return response.UploadLoanProofOfPicture{
		ID:       fileModel.ID,
		Name:     fileModel.OriginalName,
		MimeType: fileModel.MimeType,
	}, nil
}
