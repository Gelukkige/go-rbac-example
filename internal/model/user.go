package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"primaryKey"`
	Name      string         `gorm:"column:name;not null"`
	Phone     string         `gorm:"column:phone;uniqueIndex:idx_user_unique_phone,where deleted_at IS NULL"`
	Email     string         `gorm:"column:email"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Roles     []Role         `gorm:"many2many:role_users;"`
}

type UserCreateReq struct {
	Name    string   `json:"name" binding:"required"`
	Phone   string   `json:"phone" binding:"required"`
	Email   string   `json:"email"`
	RoleIDs []uint64 `json:"role_ids"`
}

type UserUpdateReq struct {
	ID      uint64   `json:"id" binding:"required"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
	Email   string   `json:"email"`
	RoleIDs []uint64 `json:"role_ids"`
}

type UserInfoResp struct {
	ID    uint64   `json:"id"`
	Name  string   `json:"name"`
	Phone string   `json:"phone"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}
