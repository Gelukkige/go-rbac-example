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
