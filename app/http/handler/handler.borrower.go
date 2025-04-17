package handler

import (
	"loan-service/pkg/helper"
	"loan-service/pkg/model/request"
	"loan-service/pkg/requesthelper"
	"loan-service/pkg/responsehelper"

	"github.com/gin-gonic/gin"
)

// GetBorrowers
//
//	@Summary	Get list of borrowers
//	@Tags		Borrower
//	@Accept		json
//	@Produce	json
//	@Param		page	query	int	false	"default is 1"
//	@Param		size	query	int	false	"default is 10"
//	@Success	200		{array}	response.GetBorrower
//	@Router		/api/v1/borrower [GET]
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

// CreateBorrower
//
//	@Summary		Create borrower
//	@Description	Borrower can only be created by user with fieldOfficer role
//	@Tags			Borrower
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.CreateBorrower	true	"Request body"
//	@Success		200		{string}	string
//	@Router			/api/v1/borrower [POST]
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

	requestedBy := c.GetString("authUser")
	newID, err := h.borrowerService.Create(&body, requestedBy)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, newID)
}

// DeleteBorrowerByID
//
//	@Summary	Delete borrower by ID
//	@Tags		Borrower
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Borrower ID"
//	@Success	200	{boolean}	true
//	@Router		/api/v1/borrower/{id} [DELETE]
func (h *Handler) DeleteBorrowerByID(c *gin.Context) {
	id := c.Param("id")
	deleted, err := h.borrowerService.DeleteByID(id)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, deleted)
}
