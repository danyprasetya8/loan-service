package handler

import (
	"loan-service/pkg/helper"
	"loan-service/pkg/model/request"
	"loan-service/pkg/requesthelper"
	"loan-service/pkg/responsehelper"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetLoans
//
//	@Summary	Get list of loans
//	@Tags		Loan
//	@Accept		json
//	@Produce	json
//	@Param		page	query	int		false	"default is 1"
//	@Param		size	query	int		false	"default is 10"
//	@Param		status	query	string	false	"filter by loan status"
//	@Success	200		{array}	response.GetLoan
//	@Router		/api/v1/loan [GET]
func (h *Handler) GetLoans(c *gin.Context) {
	var pagination request.Pagination
	if err := c.BindQuery(&pagination); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	requesthelper.SetDefaultPagination(&pagination)

	statusQuery := c.Query("status")

	list, pageRes := h.loanService.GetList(strings.TrimSpace(statusQuery), &pagination)

	responsehelper.SuccessPage(c, list, pageRes)
}

// GetLoanDetail
//
//	@Summary	Get loan detail
//	@Tags		Loan
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Loan ID"
//	@Success	200	{object}	response.GetLoanDetail
//	@Router		/api/v1/loan/{id} [GET]
func (h *Handler) GetLoanDetail(c *gin.Context) {
	loanID := c.Param("id")

	detail, err := h.loanService.GetDetail(loanID)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, detail)
}

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

// UploadLoanProofOfPicture
//
//	@Summary	Upload loan proof of picture
//	@Tags		Loan
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		id		path		string	true	"Loan ID"
//	@Param		image	formData	file	true	"Picture to upload"
//	@Success	200		{object}	response.UploadLoanProofOfPicture
//	@Router		/api/v1/loan/{id}/proof [POST]
func (h *Handler) UploadLoanProofOfPicture(c *gin.Context) {
	image, err := c.FormFile("image")

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	mimeType, err := helper.GetMimeType(image)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	if mimeType != "image/jpeg" && mimeType != "image/png" {
		responsehelper.BadRequest(c, "invalid image")
		return
	}

	loanID := c.Param("id")

	requestedBy := c.GetString("authUser")
	uploadRes, err := h.loanService.SaveProofOfPicture(image, loanID, requestedBy)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, uploadRes)
}

// ApproveLoan
//
//	@Summary		Approve loan by internal
//	@Description	User with role internal can approve a loan. fieldOfficerId is user's email, proofOfPicture is file id got from upload response
//	@Tags			Loan
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string				true	"Loan ID"
//	@Param			body	body		request.ApproveLoan	true	"Request body"
//	@Success		200		{boolean}	true
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
	body.LoanID = c.Param("id")
	approved, err := h.loanService.Approve(body, requestedBy)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, approved)
}

// InvestLoan
//
//	@Summary	Invest loan by investor
//	@Tags		Loan
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string				true	"Loan ID"
//	@Param		body	body		request.InvestLoan	true	"Request body"
//	@Success	200		{boolean}	true
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
	body.LoanID = c.Param("id")
	success, err := h.loanService.Invest(body, requestedBy)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, success)
}

// DisburseLoan
//
//	@Summary		Diburse loan by internal
//	@Description	User with role internal can approve a loan. fieldOfficerId is user's email, borrowerAgreementLetter is file id got from upload response
//	@Tags			Loan
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"Loan ID"
//	@Param			body	body		request.DisburseLoan	true	"Request body"
//	@Success		200		{boolean}	true
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
	body.LoanID = c.Param("id")
	disbursed, err := h.loanService.Disburse(body, requestedBy)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, disbursed)
}

// UploadBorrowerAgreementLetter
//
//	@Summary	Upload borrower agreement letter
//	@Tags		Loan
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		id		path		string	true	"Loan ID"
//	@Param		file	formData	file	true	"PDF to upload"
//	@Success	200		{object}	response.UploadBorrowerLetter
//	@Router		/api/v1/loan/{id}/borrower-agreement-letter [POST]
func (h *Handler) UploadBorrowerAgreementLetter(c *gin.Context) {
	pdfFile, err := c.FormFile("file")

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	mimeType, err := helper.GetMimeType(pdfFile)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	if mimeType != "application/pdf" && mimeType != "image/jpeg" {
		responsehelper.BadRequest(c, "file format must be either pdf/jpeg")
		return
	}

	loanID := c.Param("id")

	requestedBy := c.GetString("authUser")
	uploadRes, err := h.loanService.SaveBorrowerAgreementLetter(pdfFile, loanID, requestedBy)

	if err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	responsehelper.Success(c, uploadRes)
}
