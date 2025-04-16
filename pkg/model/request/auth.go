package request

import "loan-service/internal/constant"

type MockLogin struct {
	Email string            `json:"email"`
	Role  constant.UserRole `json:"role"`
}
