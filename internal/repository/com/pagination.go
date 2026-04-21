package com

// Pagination 分页
type Pagination struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func (p *Pagination) Offset() int {
	if p.Page == 0 {
		p.Page = 1
	}

	return (p.Page - 1) * p.PageSize
}

func (p *Pagination) Limit() int {
	return p.PageSize
}
