package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64
	Name      string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
