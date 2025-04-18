package loan

import (
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/mocks"
	"loan-service/pkg/model/request"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInvest_LoanNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.InvestLoan{
		LoanID: "loanId",
		Amount: 100,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(nil).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	_, err := service.Invest(req, "test@mail.com")

	assert.Equal(t, "loan not exist", err.Error())
}

func TestInvest_InvalidLoadStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.InvestLoan{
		LoanID: "loanId",
		Amount: 100,
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

	_, err := service.Invest(req, "test@mail.com")

	assert.Equal(t, "loan status must be approved", err.Error())
}

func TestInvest_ExcessAmount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.InvestLoan{
		LoanID: "loanId",
		Amount: 3000,
	}

	l := &entity.Loan{
		ID:              "loanId",
		Status:          constant.Approved,
		PrincipalAmount: 10000,
		InvestedAmount:  8000,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	_, err := service.Invest(req, "test@mail.com")

	assert.Equal(t, "remaining amount needed: 2000", err.Error())
}

func TestInvest_NewInvestor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.InvestLoan{
		LoanID: "loanId",
		Amount: 1000,
	}

	l := &entity.Loan{
		ID:              "loanId",
		Status:          constant.Approved,
		PrincipalAmount: 10000,
		InvestedAmount:  8000,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	loanInvestmentRepoMock := mocks.NewMockILoanInvestmentRepository(ctrl)
	loanInvestmentRepoMock.EXPECT().
		GetByLoanAndInvestor("loanId", "test@mail.com").
		Return(nil).
		Times(1)

	loanInvestmentRepoMock.EXPECT().
		Create(gomock.AssignableToTypeOf(&entity.LoanInvestment{})).
		DoAndReturn(func(en *entity.LoanInvestment) error {
			assert.Equal(t, int64(1000), en.Amount)
			assert.Equal(t, "test@mail.com", en.InvestorID)
			return nil
		}).
		Times(1)

	loanRepoMock.EXPECT().
		Save(gomock.AssignableToTypeOf(&entity.Loan{})).
		DoAndReturn(func(en *entity.Loan) error {
			assert.Equal(t, int64(9000), en.InvestedAmount)
			return nil
		}).
		Times(1)

	service := &Loan{
		loanRepo:           loanRepoMock,
		loanInvestmentRepo: loanInvestmentRepoMock,
	}

	success, err := service.Invest(req, "test@mail.com")

	assert.Nil(t, err)
	assert.True(t, success)
}

func TestInvest_ExistingInvestor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := request.InvestLoan{
		LoanID: "loanId",
		Amount: 1000,
	}

	l := &entity.Loan{
		ID:              "loanId",
		Status:          constant.Approved,
		PrincipalAmount: 10000,
		InvestedAmount:  8000,
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		Get("loanId").
		Return(l).
		Times(1)

	li := &entity.LoanInvestment{
		InvestorID: "test@mail.com",
		Amount:     1500,
	}
	loanInvestmentRepoMock := mocks.NewMockILoanInvestmentRepository(ctrl)
	loanInvestmentRepoMock.EXPECT().
		GetByLoanAndInvestor("loanId", "test@mail.com").
		Return(li).
		Times(1)

	loanInvestmentRepoMock.EXPECT().
		Save(gomock.AssignableToTypeOf(&entity.LoanInvestment{})).
		DoAndReturn(func(en *entity.LoanInvestment) error {
			assert.Equal(t, int64(2500), en.Amount)
			assert.Equal(t, "test@mail.com", en.InvestorID)
			return nil
		}).
		Times(1)

	loanRepoMock.EXPECT().
		Save(gomock.AssignableToTypeOf(&entity.Loan{})).
		DoAndReturn(func(en *entity.Loan) error {
			assert.Equal(t, int64(9000), en.InvestedAmount)
			return nil
		}).
		Times(1)

	service := &Loan{
		loanRepo:           loanRepoMock,
		loanInvestmentRepo: loanInvestmentRepoMock,
	}

	success, err := service.Invest(req, "test@mail.com")

	assert.Nil(t, err)
	assert.True(t, success)
}
