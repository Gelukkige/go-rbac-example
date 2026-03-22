package model

const (
	DefaultPageNum  = 1
	DefaultPageSize = 10
)

type Page struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

type DeleteIDs struct {
	IDs []uint64 `json:"ids" binding:"required"`
}
