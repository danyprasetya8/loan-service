package loan

import (
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/mocks"
	"loan-service/pkg/helper"
	"loan-service/pkg/model/request"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	nowFormatted := helper.FormatDate(now)

	loans := []entity.Loan{
		{
			ID:              "id",
			BorrowerID:      "borrower1",
			Status:          constant.Approved,
			PrincipalAmount: 1000,
			InvestedAmount:  800,
			Rate:            5,
			ROI:             5,
			Audit: entity.Audit{
				CreatedAt: now,
				CreatedBy: "test@mail.com",
			},
		},
		{
			ID:              "id2",
			BorrowerID:      "borrower2",
			Status:          constant.Approved,
			PrincipalAmount: 2000,
			InvestedAmount:  1200,
			Rate:            8,
			ROI:             8,
			Audit: entity.Audit{
				CreatedAt: now,
				CreatedBy: "test@mail.com",
			},
		},
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		GetList(constant.Approved, 1, 2).
		Return(loans).
		Times(1)

	loanRepoMock.EXPECT().
		Count(constant.Approved).
		Return(int64(11))

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	list, pageRes := service.GetList(string(constant.Approved), &request.Pagination{
		Page: 1,
		Size: 2,
	})

	assert.Equal(t, 2, len(list))
	assert.Equal(t, "id", list[0].ID)
	assert.Equal(t, nowFormatted, list[0].CreatedAt)
	assert.NotNil(t, pageRes)
	assert.Equal(t, 1, pageRes.Page)
	assert.Equal(t, 2, pageRes.Size)
	assert.Equal(t, int64(11), pageRes.TotalData)
	assert.Equal(t, 6, pageRes.TotalPage)
}

func TestGetDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := &entity.Loan{
		ID: "loanId",
		LoanApproval: entity.LoanApproval{
			FieldOfficerID: "officer@mail.com",
			ProofOfPicture: "proofId",
		},
		LoanInvestment: []entity.LoanInvestment{
			{InvestorID: "investor@mail.com", Amount: 1000},
		},
		LoanDisbursement: entity.LoanDisbursement{
			FieldOfficerID:          "officer@mail.com",
			BorrowerAgreementLetter: "letterId",
		},
	}

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		GetDetail("loanId").
		Return(l).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	detail, err := service.GetDetail("loanId")

	assert.Nil(t, err)
	assert.NotNil(t, detail)
	assert.Equal(t, "loanId", detail.ID)
	assert.Equal(t, "officer@mail.com", detail.Approval.FieldOfficerID)
	assert.Equal(t, "proofId", detail.Approval.ProofOfPicture)
	assert.Equal(t, "officer@mail.com", detail.Disbursement.FieldOfficerID)
	assert.Equal(t, "letterId", detail.Disbursement.BorrowerAgreementLetter)
	assert.Equal(t, 1, len(detail.Investors))
	assert.Equal(t, "investor@mail.com", detail.Investors[0].InvestorID)
}

func TestGetDetail_LoanNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	loanRepoMock := mocks.NewMockILoanRepository(ctrl)
	loanRepoMock.EXPECT().
		GetDetail("loanId").
		Return(nil).
		Times(1)

	service := &Loan{
		loanRepo: loanRepoMock,
	}

	_, err := service.GetDetail("loanId")

	assert.Equal(t, "loan not exist", err.Error())
}
