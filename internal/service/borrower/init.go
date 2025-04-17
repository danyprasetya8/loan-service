package borrower

import (
	"loan-service/internal/entity"
	repo "loan-service/internal/repository/borrower"
	"loan-service/pkg/helper"
	"loan-service/pkg/model/request"
	"loan-service/pkg/model/response"
	"loan-service/pkg/responsehelper"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

type IBorrowerService interface {
	Create(req *request.CreateBorrower, requestedBy string) (id string, err error)
	GetList(page *request.Pagination) (list []response.GetBorrower, pageRes *response.Pagination)
	DeleteByID(id string) (deleted bool, err error)
}

type Borrower struct {
	repo repo.IBorrowerRepository
}

func New(repo repo.IBorrowerRepository) IBorrowerService {
	return &Borrower{repo}
}

func (s *Borrower) Create(req *request.CreateBorrower, requestedBy string) (id string, err error) {
	newBorrower := &entity.Borrower{
		ID:   uuid.New().String(),
		Name: req.Name,
		Audit: entity.Audit{
			CreatedBy: requestedBy,
			UpdatedBy: requestedBy,
		},
	}
	err = s.repo.Create(newBorrower)
	if err != nil {
		log.Errorf("Failed to create borrower: %s", err.Error())
	}
	return newBorrower.ID, err
}

func (s *Borrower) GetList(page *request.Pagination) (list []response.GetBorrower, pageRes *response.Pagination) {
	borrowers := s.repo.GetList(page.Page, page.Size)
	total := s.repo.Count()

	for _, borrower := range borrowers {
		list = append(list, response.GetBorrower{
			ID:        borrower.ID,
			Name:      borrower.Name,
			CreatedBy: borrower.CreatedBy,
			CreatedAt: helper.FormatDate(borrower.CreatedAt),
		})
	}

	return list, responsehelper.ToPagination(page, total)
}

func (s *Borrower) DeleteByID(id string) (deleted bool, err error) {
	deleted, err = s.repo.Delete(id)
	if err != nil {
		log.Errorf("Failed to delete borrower: %s", err.Error())
	}
	return
}
