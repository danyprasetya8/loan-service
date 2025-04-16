package response

type Base[T any] struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       T           `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type Pagination struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalPage int   `json:"totalPage"`
	TotalData int64 `json:"totalData"`
}
