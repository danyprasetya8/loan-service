package loan

import (
	"errors"
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/internal/service/file"
	"loan-service/mocks"
	"mime/multipart"
	"net/textproto"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSaveProofOfPicture_HasError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(nil).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	m := &multipart.FileHeader{
		Filename: "empty.txt",
		Header:   textproto.MIMEHeader{},
		Size:     0,
	}
	_, err := service.SaveProofOfPicture(m, "loanId", "test@mail.com")
	assert.Equal(t, "loan not exist", err.Error())
}

func TestSaveProofOfPicture_InvalidLoanStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := &entity.Loan{
		ID:     "loanId",
		Status: constant.Invested,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	m := &multipart.FileHeader{
		Filename: "empty.txt",
		Header:   textproto.MIMEHeader{},
		Size:     0,
	}
	_, err := service.SaveProofOfPicture(m, "loanId", "test@mail.com")
	assert.Equal(t, "loan status must be proposed", err.Error())
}

func TestSaveProofOfPicture_ErrorFromFileService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := &entity.Loan{
		ID:     "loanId",
		Status: constant.Proposed,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	m := &multipart.FileHeader{
		Filename: "empty.txt",
		Header:   textproto.MIMEHeader{},
		Size:     0,
	}

	fileServiceMock := mocks.NewMockIFileService(ctrl)
	fileServiceMock.EXPECT().
		Save(m, constant.ProofOfPicture, filepath.Join("proofOfPicture", "loanId"), "test@mail.com").
		Return(nil, errors.New("some error")).
		Times(1)

	service := &Loan{
		loanRepo:    loanRepoMock,
		fileService: fileServiceMock,
	}

	_, err := service.SaveProofOfPicture(m, "loanId", "test@mail.com")
	assert.Equal(t, "some error", err.Error())
}

func TestSaveProofOfPicture(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := &entity.Loan{
		ID:     "loanId",
		Status: constant.Proposed,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	m := &multipart.FileHeader{
		Filename: "empty.txt",
		Header:   textproto.MIMEHeader{},
		Size:     0,
	}

	fm := &file.Model{
		ID:           "fileId",
		OriginalName: "name",
		Path:         "path/to/file.pdf",
		MimeType:     "application/pdf",
		Type:         constant.ProofOfPicture,
	}

	fileServiceMock := mocks.NewMockIFileService(ctrl)
	fileServiceMock.EXPECT().
		Save(m, constant.ProofOfPicture, filepath.Join("proofOfPicture", "loanId"), "test@mail.com").
		Return(fm, nil).
		Times(1)

	service := &Loan{
		loanRepo:    loanRepoMock,
		fileService: fileServiceMock,
	}

	res, err := service.SaveProofOfPicture(m, "loanId", "test@mail.com")
	assert.Nil(t, err)
	assert.Equal(t, "fileId", res.ID)
	assert.Equal(t, "application/pdf", res.MimeType)
	assert.Equal(t, "name", res.Name)
}
