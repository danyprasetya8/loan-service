package response

type GetUser struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}
