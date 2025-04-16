package handler

import (
	"loan-service/pkg/helper"
	"loan-service/pkg/model/request"
	"loan-service/pkg/responsehelper"

	"github.com/gin-gonic/gin"
)

// MockLogin
//
//	@Summary	Mock login using email and user role
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		body	body		request.MockLogin	true	"Request body"
//	@Success	200		{string}	string
//	@Router		/api/v1/auth/mock-login [POST]
func (h *Handler) MockLogin(c *gin.Context) {
	var body request.MockLogin
	if err := c.BindJSON(&body); err != nil {
		responsehelper.BadRequest(c, err.Error())
		return
	}

	if !helper.IsValidEmail(body.Email) {
		responsehelper.BadRequest(c, "email must be valid")
		return
	}

	token, err := h.authService.MockLogin(&body)

	if err != nil {
		responsehelper.Unauthenticated(c)
		return
	}

	c.SetCookie("token", token, 3600*24*30, "/", "", false, true)
	responsehelper.Success(c, token)
}
