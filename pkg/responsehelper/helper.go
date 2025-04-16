package responsehelper

import (
	"loan-service/pkg/model/request"
	"loan-service/pkg/model/response"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, &response.Base[T]{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func SuccessPage[T any](c *gin.Context, data T, pagination *response.Pagination) {
	c.JSON(http.StatusOK, &response.Base[T]{
		Code:       http.StatusOK,
		Message:    "success",
		Data:       data,
		Pagination: pagination,
	})
}

func BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, &response.Error{
		Code:    http.StatusBadRequest,
		Message: "error",
		Error:   msg,
	})
}

func Unauthenticated(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, &response.Error{
		Code:    http.StatusUnauthorized,
		Message: "error",
		Error:   "unauthenticated",
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusForbidden, &response.Error{
		Code:    http.StatusForbidden,
		Message: "error",
		Error:   "forbidden",
	})
}

func ToPagination(request *request.Pagination, total int64) *response.Pagination {
	return &response.Pagination{
		Page:      request.Page,
		Size:      request.Size,
		TotalPage: int(math.Ceil(float64(total) / float64(request.Size))),
		TotalData: total,
	}
}
