package model

type DeleteIDs struct {
	IDs []uint64 `json:"ids" binding:"required"`
}
