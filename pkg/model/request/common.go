package request

type Pagination struct {
	Page int `form:"page"`
	Size int `form:"size"`
}
