package model

type Page struct {
	PageNum  int `form:"page_num" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=10,max=100"`
}

type DeleteIDs struct {
	IDs []uint64 `json:"ids" binding:"required"`
}
