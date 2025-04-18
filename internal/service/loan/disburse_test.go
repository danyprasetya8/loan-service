package loan

import (
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/internal/service/file"
	"loan-service/mocks"
	"loan-service/pkg/model/request"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDisburse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.DisburseLoan{
		LoanID:                  "loanId",
		FieldOfficerID:          "officer@mail.com",
		BorrowerAgreementLetter: "letterId",
	}

	l := &entity.Loan{
		ID:     "loanId",
		Status: constant.Invested,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	u := &entity.User{
		Email: "officer@mail.com",
		Role:  constant.FieldOfficer,
	}
	userRepoMock := mocks.NewMockIUserRepository(ctrl)
	userRepoMock.EXPECT().
		Get("officer@mail.com").
		Return(u).
		Times(1)

	f := &file.Model{
		ID:   "letterId",
		Path: "path/to/file.pdf",
	}
	fileServiceMock := mocks.NewMockIFileService(ctrl)
	fileServiceMock.EXPECT().
		Find("letterId").
		Return(f).
		Times(1)

	fileServiceMock.EXPECT().
		IsExist(f.Path).
		Return(true).
		Times(1)

	loanDisbursementRepoMock := mocks.NewMockILoanDisbursementRepository(ctrl)
	loanDisbursementRepoMock.EXPECT().
		Create(gomock.AssignableToTypeOf(&entity.LoanDisbursement{})).
		DoAndReturn(func(en *entity.LoanDisbursement) error {
			assert.Equal(t, "loanId", en.LoanID)
			assert.Equal(t, "officer@mail.com", en.FieldOfficerID)
			assert.Equal(t, "letterId", en.BorrowerAgreementLetter)
			assert.Equal(t, "test@mail.com", en.Audit.CreatedBy)
			assert.Equal(t, "test@mail.com", en.Audit.UpdatedBy)
			assert.NotEmpty(t, en.ID)
			return nil
		})

	service := &Loan{
		loanRepo:             loanRepoMock,
		userRepo:             userRepoMock,
		fileService:          fileServiceMock,
		loanDisbursementRepo: loanDisbursementRepoMock,
	}

	loanRepoMock.EXPECT().
		Save(gomock.AssignableToTypeOf(&entity.Loan{})).
		DoAndReturn(func(en *entity.Loan) error {
			assert.Equal(t, constant.Disbursed, en.Status)
			return nil
		}).
		Times(1)

	success, err := service.Disburse(req, "test@mail.com")
	assert.Nil(t, err)
	assert.True(t, success)
}

func TestDisburse_LoanNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.DisburseLoan{
		LoanID:                  "loanId",
		FieldOfficerID:          "officer@mail.com",
		BorrowerAgreementLetter: "letterId",
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(nil)

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	_, err := service.Disburse(req, "test@mail.com")
	assert.Equal(t, "loan not exist", err.Error())
}

func TestDisburse_InvalidLoanStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.DisburseLoan{
		LoanID:                  "loanId",
		FieldOfficerID:          "officer@mail.com",
		BorrowerAgreementLetter: "letterId",
	}

	l := &entity.Loan{
		ID:     "loanId",
		Status: constant.Proposed,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	_, err := service.Disburse(req, "test@mail.com")
	assert.Equal(t, "loan status must be invested", err.Error())
}

func TestDisburse_FieldOfficerNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.DisburseLoan{
		LoanID:                  "loanId",
		FieldOfficerID:          "officer@mail.com",
		BorrowerAgreementLetter: "letterId",
	}

	l := &entity.Loan{
		ID:     "loanId",
		Status: constant.Invested,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	userRepoMock := mocks.NewMockIUserRepository(ctrl)
	userRepoMock.EXPECT().
		Get("officer@mail.com").
		Return(nil).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
		userRepo: userRepoMock,
	}

	_, err := service.Disburse(req, "test@mail.com")
	assert.Equal(t, "field officer not exist", err.Error())
}

func TestDisburse_UserNotFieldOfficer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.DisburseLoan{
		LoanID:                  "loanId",
		FieldOfficerID:          "officer@mail.com",
		BorrowerAgreementLetter: "letterId",
	}

	l := &entity.Loan{
		ID:     "loanId",
		Status: constant.Invested,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	u := &entity.User{
		Email: "officer@mail.com",
		Role:  constant.Internal,
	}
	userRepoMock := mocks.NewMockIUserRepository(ctrl)
	userRepoMock.EXPECT().
		Get("officer@mail.com").
		Return(u).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
		userRepo: userRepoMock,
	}

	_, err := service.Disburse(req, "test@mail.com")
	assert.Equal(t, "invalid user role", err.Error())
}

func TestDisburse_FileNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.DisburseLoan{
		LoanID:                  "loanId",
		FieldOfficerID:          "officer@mail.com",
		BorrowerAgreementLetter: "letterId",
	}

	l := &entity.Loan{
		ID:     "loanId",
		Status: constant.Invested,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	u := &entity.User{
		Email: "officer@mail.com",
		Role:  constant.FieldOfficer,
	}
	userRepoMock := mocks.NewMockIUserRepository(ctrl)
	userRepoMock.EXPECT().
		Get("officer@mail.com").
		Return(u).
		Times(1)

	fileServiceMock := mocks.NewMockIFileService(ctrl)
	fileServiceMock.EXPECT().
		Find("letterId").
		Return(nil).
		Times(1)

	service := &Loan{
		loanRepo:    loanRepoMock,
		userRepo:    userRepoMock,
		fileService: fileServiceMock,
	}

	_, err := service.Disburse(req, "test@mail.com")
	assert.Equal(t, "agreement letter not exist", err.Error())
}

func TestDisburse_FileNotExistInDisk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.DisburseLoan{
		LoanID:                  "loanId",
		FieldOfficerID:          "officer@mail.com",
		BorrowerAgreementLetter: "letterId",
	}

	l := &entity.Loan{
		ID:     "loanId",
		Status: constant.Invested,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	u := &entity.User{
		Email: "officer@mail.com",
		Role:  constant.FieldOfficer,
	}
	userRepoMock := mocks.NewMockIUserRepository(ctrl)
	userRepoMock.EXPECT().
		Get("officer@mail.com").
		Return(u).
		Times(1)

	f := &file.Model{
		ID:   "letterId",
		Path: "path/to/file.pdf",
	}
	fileServiceMock := mocks.NewMockIFileService(ctrl)
	fileServiceMock.EXPECT().
		Find("letterId").
		Return(f).
		Times(1)

	fileServiceMock.EXPECT().
		IsExist(f.Path).
		Return(false).
		Times(1)

	service := &Loan{
		loanRepo:    loanRepoMock,
		userRepo:    userRepoMock,
		fileService: fileServiceMock,
	}

	_, err := service.Disburse(req, "test@mail.com")
	assert.Equal(t, "agreement letter not exist", err.Error())
}
