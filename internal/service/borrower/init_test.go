package borrower

import (
	"errors"
	"loan-service/internal/entity"
	"loan-service/mocks"
	"loan-service/pkg/model/request"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreate_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	borrowerRepoMock := mocks.NewMockIBorrowerRepository(ctrl)
	borrowerRepoMock.EXPECT().
		Create(gomock.AssignableToTypeOf(&entity.Borrower{})).
		DoAndReturn(func(en *entity.Borrower) error {
			en.ID = "mock-id"
			return nil
		})

	service := &Borrower{
		repo: borrowerRepoMock,
	}

	id, err := service.Create(&request.CreateBorrower{
		Name: "borrower1",
	}, "test@mail.com")

	assert.Nil(t, err)
	assert.Equal(t, "mock-id", id)
}

func TestCreate_HasError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedErr := errors.New("should be this error")

	borrowerRepoMock := mocks.NewMockIBorrowerRepository(ctrl)
	borrowerRepoMock.EXPECT().
		Create(gomock.AssignableToTypeOf(&entity.Borrower{})).
		Return(expectedErr)

	service := &Borrower{
		repo: borrowerRepoMock,
	}

	_, err := service.Create(&request.CreateBorrower{
		Name: "borrower1",
	}, "test@mail.com")

	assert.Equal(t, expectedErr, err)
}

func TestGetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	borrowers := []entity.Borrower{
		{ID: "id1", Name: "borrower1"},
		{ID: "id2", Name: "borrower2"},
	}

	borrowerRepoMock := mocks.NewMockIBorrowerRepository(ctrl)
	borrowerRepoMock.EXPECT().
		GetList(1, 2).
		Return(borrowers)
	borrowerRepoMock.EXPECT().
		Count().
		Return(int64(11))

	service := &Borrower{
		repo: borrowerRepoMock,
	}

	list, pageRes := service.GetList(&request.Pagination{
		Page: 1,
		Size: 2,
	})

	assert.Equal(t, 2, len(list))
	assert.Equal(t, "id1", list[0].ID)
	assert.Equal(t, "id2", list[1].ID)
	assert.NotNil(t, pageRes)
	assert.Equal(t, 1, pageRes.Page)
	assert.Equal(t, 2, pageRes.Size)
	assert.Equal(t, int64(11), pageRes.TotalData)
	assert.Equal(t, 6, pageRes.TotalPage)
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	borrowerRepoMock := mocks.NewMockIBorrowerRepository(ctrl)
	borrowerRepoMock.EXPECT().
		Delete("id").
		Return(true, nil)

	service := &Borrower{
		repo: borrowerRepoMock,
	}

	deleted, err := service.DeleteByID("id")
	assert.True(t, deleted)
	assert.Nil(t, err)
}

func TestDelete_HasError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedErr := errors.New("should be this error")

	borrowerRepoMock := mocks.NewMockIBorrowerRepository(ctrl)
	borrowerRepoMock.EXPECT().
		Delete("id").
		Return(false, expectedErr)

	service := &Borrower{
		repo: borrowerRepoMock,
	}

	deleted, err := service.DeleteByID("id")
	assert.False(t, deleted)
	assert.Equal(t, expectedErr, err)
}
