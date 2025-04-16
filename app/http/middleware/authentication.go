package middleware

import (
	"loan-service/internal/constant"
	"loan-service/internal/service/auth"
	"loan-service/pkg/helper"
	"loan-service/pkg/responsehelper"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	authService auth.IAuthService
}

func New(authService auth.IAuthService) *Middleware {
	return &Middleware{authService}
}

func (m *Middleware) Authenticate(c *gin.Context) {
	tokenStr, err := c.Cookie("token")
	if err != nil {
		responsehelper.Unauthenticated(c)
		c.Abort()
		return
	}

	email, err := m.authService.ParseToken(tokenStr)
	if err != nil {
		responsehelper.Unauthenticated(c)
		c.Abort()
		return
	}

	c.Set("authUser", email)
	c.Next()
}

func (m *Middleware) Authorize(role constant.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString("authUser")

		if helper.IsBlank(email) {
			responsehelper.Unauthorized(c)
			c.Abort()
			return
		}

		userRole, err := m.authService.GetUserRole(email)

		if err != nil {
			responsehelper.Unauthorized(c)
			c.Abort()
			return
		}

		if userRole != role {
			responsehelper.Unauthorized(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
