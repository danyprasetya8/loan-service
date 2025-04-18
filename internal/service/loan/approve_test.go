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

func TestApprove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.ApproveLoan{
		LoanID:         "loanId",
		FieldOfficerID: "officer@mail.com",
		ProofOfPicture: "proofId",
	}

	l := &entity.Loan{
		ID:         "loanId",
		BorrowerID: "borrowerId",
		Status:     constant.Proposed,
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
		ID:   "proofId",
		Path: "/path/to/file.pdf",
	}

	fileServiceMock := mocks.NewMockIFileService(ctrl)
	fileServiceMock.EXPECT().
		Find("proofId").
		Return(f).
		Times(1)
	fileServiceMock.EXPECT().
		IsExist(f.Path).
		Return(true).
		Times(1)

	loanApprovalRepoMock := mocks.NewMockILoanApprovalRepository(ctrl)
	loanApprovalRepoMock.EXPECT().
		Create(gomock.AssignableToTypeOf(&entity.LoanApproval{})).
		DoAndReturn(func(en *entity.LoanApproval) error {
			assert.Equal(t, "loanId", en.LoanID)
			assert.Equal(t, "officer@mail.com", en.FieldOfficerID)
			assert.Equal(t, "proofId", en.ProofOfPicture)
			assert.Equal(t, "test@mail.com", en.Audit.CreatedBy)
			assert.Equal(t, "test@mail.com", en.Audit.UpdatedBy)
			assert.NotEmpty(t, en.ID)
			return nil
		})

	loanRepoMock.EXPECT().
		Save(gomock.AssignableToTypeOf(&entity.Loan{})).
		DoAndReturn(func(en *entity.Loan) error {
			assert.Equal(t, constant.Approved, en.Status)
			return nil
		}).
		Times(1)

	service := &Loan{
		loanRepo:         loanRepoMock,
		userRepo:         userRepoMock,
		fileService:      fileServiceMock,
		loanApprovalRepo: loanApprovalRepoMock,
	}

	success, err := service.Approve(req, "test@mail.com")
	assert.Nil(t, err)
	assert.True(t, success)
}

func TestApprove_LoanNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.ApproveLoan{
		LoanID:         "loanId",
		FieldOfficerID: "officer@mail.com",
		ProofOfPicture: "proofId",
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(nil).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	_, err := service.Approve(req, "test@mail.com")
	assert.Equal(t, "loan not exist", err.Error())
}

func TestApprove_InvalidLoanStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.ApproveLoan{
		LoanID:         "loanId",
		FieldOfficerID: "officer@mail.com",
		ProofOfPicture: "proofId",
	}

	l := &entity.Loan{
		ID:         "loanId",
		BorrowerID: "borrowerId",
		Status:     constant.Invested,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	_, err := service.Approve(req, "test@mail.com")
	assert.Equal(t, "loan status must be proposed", err.Error())
}

func TestApprove_FieldOfficerNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.ApproveLoan{
		LoanID:         "loanId",
		FieldOfficerID: "officer@mail.com",
		ProofOfPicture: "proofId",
	}

	l := &entity.Loan{
		ID:         "loanId",
		BorrowerID: "borrowerId",
		Status:     constant.Proposed,
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

	_, err := service.Approve(req, "test@mail.com")
	assert.Equal(t, "field officer not exist", err.Error())
}

func TestApprove_UserIsNotFieldOfficer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.ApproveLoan{
		LoanID:         "loanId",
		FieldOfficerID: "officer@mail.com",
		ProofOfPicture: "proofId",
	}

	l := &entity.Loan{
		ID:         "loanId",
		BorrowerID: "borrowerId",
		Status:     constant.Proposed,
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

	_, err := service.Approve(req, "test@mail.com")
	assert.Equal(t, "invalid user role", err.Error())
}

func TestApprove_FileNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.ApproveLoan{
		LoanID:         "loanId",
		FieldOfficerID: "officer@mail.com",
		ProofOfPicture: "proofId",
	}

	l := &entity.Loan{
		ID:         "loanId",
		BorrowerID: "borrowerId",
		Status:     constant.Proposed,
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
		Find("proofId").
		Return(nil).
		Times(1)

	service := &Loan{
		loanRepo:    loanRepoMock,
		userRepo:    userRepoMock,
		fileService: fileServiceMock,
	}

	_, err := service.Approve(req, "test@mail.com")
	assert.Equal(t, "proof of picture not exist", err.Error())
}

func TestApprove_FileNotExistInDisk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.ApproveLoan{
		LoanID:         "loanId",
		FieldOfficerID: "officer@mail.com",
		ProofOfPicture: "proofId",
	}

	l := &entity.Loan{
		ID:         "loanId",
		BorrowerID: "borrowerId",
		Status:     constant.Proposed,
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
		ID:   "proofId",
		Path: "/path/to/file.pdf",
	}

	fileServiceMock := mocks.NewMockIFileService(ctrl)
	fileServiceMock.EXPECT().
		Find("proofId").
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

	_, err := service.Approve(req, "test@mail.com")
	assert.Equal(t, "proof of picture not exist", err.Error())
}
