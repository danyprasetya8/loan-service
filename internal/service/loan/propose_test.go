package loan

import (
	"errors"
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/mocks"
	"loan-service/pkg/model/request"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPropose(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.ProposeLoan{
		BorrowerID:      "borrowerId",
		PrincipalAmount: 1000,
		Rate:            5,
		ROI:             5,
	}

	br := &entity.Borrower{
		ID:   "borrowerId",
		Name: "borrowerName",
	}

	borrowerRepoMock := mocks.NewMockIBorrowerRepository(ctrl)
	borrowerRepoMock.EXPECT().
		GetByID("borrowerId").
		Return(br).
		Times(1)

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Create(&entity.Loan{
			BorrowerID:      "borrowerId",
			Status:          constant.Proposed,
			PrincipalAmount: 1000,
			InvestedAmount:  0,
			Rate:            5,
			ROI:             5,
			Audit: entity.Audit{
				CreatedBy: "test@mail.com",
				UpdatedBy: "test@mail.com",
			},
		}).
		DoAndReturn(func(en *entity.Loan) error {
			en.ID = "mockId"
			return nil
		}).
		Times(1)

	service := &Loan{
		borrowerRepo: borrowerRepoMock,
		loanRepo:     loanRepoMock,
	}

	id, err := service.Propose(req, "test@mail.com")
	assert.Nil(t, err)
	assert.Equal(t, id, "mockId")
}

func TestPropose_BorrowerNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.ProposeLoan{
		BorrowerID:      "borrowerId",
		PrincipalAmount: 1000,
		Rate:            5,
		ROI:             5,
	}

	borrowerRepoMock := mocks.NewMockIBorrowerRepository(ctrl)
	borrowerRepoMock.EXPECT().
		GetByID("borrowerId").
		Return(nil).
		Times(1)

	service := &Loan{
		borrowerRepo: borrowerRepoMock,
	}

	_, err := service.Propose(req, "test@mail.com")
	assert.Equal(t, "borrower not exist", err.Error())
}

func TestPropose_HasError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.ProposeLoan{
		BorrowerID:      "borrowerId",
		PrincipalAmount: 1000,
		Rate:            5,
		ROI:             5,
	}

	br := &entity.Borrower{
		ID:   "borrowerId",
		Name: "borrowerName",
	}

	borrowerRepoMock := mocks.NewMockIBorrowerRepository(ctrl)
	borrowerRepoMock.EXPECT().
		GetByID("borrowerId").
		Return(br).
		Times(1)

	expectedErr := errors.New("should be this error")

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Create(&entity.Loan{
			BorrowerID:      "borrowerId",
			Status:          constant.Proposed,
			PrincipalAmount: 1000,
			InvestedAmount:  0,
			Rate:            5,
			ROI:             5,
			Audit: entity.Audit{
				CreatedBy: "test@mail.com",
				UpdatedBy: "test@mail.com",
			},
		}).
		Return(expectedErr).
		Times(1)

	service := &Loan{
		borrowerRepo: borrowerRepoMock,
		loanRepo:     loanRepoMock,
	}

	_, err := service.Propose(req, "test@mail.com")
	assert.Equal(t, expectedErr, err)
}
