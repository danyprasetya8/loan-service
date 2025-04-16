package requesthelper

import "loan-service/pkg/model/request"

func SetDefaultPagination(pagination *request.Pagination) {
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	if pagination.Size == 0 {
		pagination.Size = 10
	}
}
