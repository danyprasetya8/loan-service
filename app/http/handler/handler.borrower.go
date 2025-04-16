package handler

import (
	"loan-service/pkg/helper"
	"loan-service/pkg/model/request"
	"loan-service/pkg/requesthelper"
	"loan-service/pkg/responsehelper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBorrowers(c *gin.Context) {
	var pagination request.Pagination
	if err := c.BindQuery(&pagination); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	requesthelper.SetDefaultPagination(&pagination)

	list, pageRes := h.borrowerService.GetList(&pagination)

	responsehelper.SuccessPage(c, list, pageRes)
}

func (h *Handler) CreateBorrower(c *gin.Context) {
	var body request.CreateBorrower
	if err := c.BindJSON(&body); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	if helper.IsBlank(body.Name) {
		responsehelper.BadRequest(c, "name must not blank")
		return
	}

	newID, err := h.borrowerService.Create(&body)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, newID)
}

func (h *Handler) DeleteBorrowerByID(c *gin.Context) {
	id := c.Param("id")
	deleted, err := h.borrowerService.DeleteByID(id)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, deleted)
}
