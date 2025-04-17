package handler

import (
	"loan-service/pkg/helper"
	"loan-service/pkg/model/request"
	"loan-service/pkg/responsehelper"

	"github.com/gin-gonic/gin"
)

// ProposeLoan
//
//	@Summary		Propose loan by field officer
//	@Description	User with role fieldOfficer can propose a loan for a borrower
//	@Tags			Loan
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.ProposeLoan	true	"Request body"
//	@Success		200		{string}	string
//	@Router			/api/v1/loan [POST]
func (h *Handler) ProposeLoan(c *gin.Context) {
	var body request.ProposeLoan
	if err := c.BindJSON(&body); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	if helper.IsBlank(body.BorrowerID) {
		responsehelper.BadRequest(c, "borrowerId must not blank")
		return
	}

	if body.PrincipalAmount <= 0 {
		responsehelper.BadRequest(c, "principalAmount must be greater than 0")
		return
	}

	if body.Rate <= 0 || body.ROI <= 0 {
		responsehelper.BadRequest(c, "rate/roi must be greater than 0")
		return
	}

	requestedBy := c.GetString("authUser")
	newID, err := h.loanService.Propose(body, requestedBy)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, newID)
}

// ApproveLoan
//
//	@Summary		Approve loan by internal
//	@Description	User with role internal can approve a loan. fieldOfficerId is user's email, proofOfPicture is path got from upload response
//	@Tags			Loan
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.ApproveLoan	true	"Request body"
//	@Success		200		{string}	string
//	@Router			/api/v1/loan/{id}/_approve [POST]
func (h *Handler) ApproveLoan(c *gin.Context) {
	var body request.ApproveLoan
	if err := c.BindJSON(&body); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	if helper.IsBlank(body.FieldOfficerID) {
		responsehelper.BadRequest(c, "fieldOfficerId must not blank")
		return
	}

	if helper.IsBlank(body.ProofOfPicture) {
		responsehelper.BadRequest(c, "proofOfPicture must not blank")
		return
	}

	requestedBy := c.GetString("authUser")
	newID, err := h.loanService.Approve(body, requestedBy)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, newID)
}

// InvestLoan
//
//	@Summary	Invest loan by investor
//	@Tags		Loan
//	@Accept		json
//	@Produce	json
//	@Param		body	body		request.InvestLoan	true	"Request body"
//	@Success	200		{string}	string
//	@Router		/api/v1/loan/{id}/_invest [POST]
func (h *Handler) InvestLoan(c *gin.Context) {
	var body request.InvestLoan
	if err := c.BindJSON(&body); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	if body.Amount <= 0 {
		responsehelper.BadRequest(c, "amount must be greater than 0")
		return
	}

	requestedBy := c.GetString("authUser")
	newID, err := h.loanService.Invest(body, requestedBy)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, newID)
}

// DisburseLoan
//
//	@Summary		Invest loan by internal
//	@Description	User with role internal can approve a loan. fieldOfficerId is user's email, borrowerAgreementLetter is path got from upload response
//	@Tags			Loan
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.DisburseLoan	true	"Request body"
//	@Success		200		{string}	string
//	@Router			/api/v1/loan/{id}/_disburse [POST]
func (h *Handler) DisburseLoan(c *gin.Context) {
	var body request.DisburseLoan
	if err := c.BindJSON(&body); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	if helper.IsBlank(body.FieldOfficerID) {
		responsehelper.BadRequest(c, "fieldOfficerId must not blank")
		return
	}

	if helper.IsBlank(body.BorrowerAgreementLetter) {
		responsehelper.BadRequest(c, "borrowerAgreementLetter must not blank")
		return
	}

	requestedBy := c.GetString("authUser")
	newID, err := h.loanService.Disburse(body, requestedBy)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, newID)
}
