package basic

// Pagination 分页
type Pagination struct {
	Page     int32 `json:"page" form:"page"`
	PageSize int32 `json:"page_size" form:"page_size"`
}
