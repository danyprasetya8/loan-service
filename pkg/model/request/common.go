package request

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}
