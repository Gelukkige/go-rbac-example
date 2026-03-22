package model

type Permission struct {
	ID      uint64   `gorm:"primaryKey"`
	Page    string   `gorm:"column:page;not null;index:idx_permission_page_action"`
	Action  string   `gorm:"column:action;not null;index:idx_permission_page_action"`
	Columns []string `gorm:"column:columns;type:jsonb;serializer:json"`
	Roles   []Role   `gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE;"`
}

type PermissionReq struct {
	Page    string   `json:"page" binding:"required"`
	Action  string   `json:"action" binding:"required"`
	Columns []string `json:"columns"`
}

type PermissionUpdateReq struct {
	ID      uint64   `json:"id" binding:"required"`
	Page    string   `json:"page"`
	Action  string   `json:"action"`
	Columns []string `json:"columns"`
}

// type PermissionInfoResp struct {
// 	ID      uint64   `json:"id"`
// 	Page    string   `json:"page"`
// 	Action  string   `json:"action"`
// 	Columns []string `json:"columns"`
// }
