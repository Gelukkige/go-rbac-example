package model

type Role struct {
	ID          uint64       `gorm:"primaryKey"`
	Name        string       `gorm:"column:name;not null"`
	Desc        string       `gorm:"column:desc"`
	Users       []User       `gorm:"many2many:role_users;"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type RoleCreateReq struct {
	Name          string   `json:"name" binding:"required"`
	Desc          string   `json:"desc"`
	UserIDs       []uint64 `json:"user_ids"`
	PermissionIDs []uint64 `json:"permission_ids"`
}

type RoleUpdateReq struct {
	ID            uint64   `json:"id" binding:"required"`
	Name          string   `json:"name"`
	Desc          string   `json:"desc"`
	UserIDs       []uint64 `json:"user_ids"`
	PermissionIDs []uint64 `json:"permission_ids"`
}

type RoleInfoResp struct {
	ID          uint64   `json:"id"`
	Name        string   `json:"name"`
	Desc        string   `json:"desc"`
	Users       []string `json:"users"`
	Permissions []string `json:"permissions"`
}
