package model

type Role struct {
	ID          uint64       `gorm:"primaryKey"`
	Name        string       `gorm:"column:name;not null"`
	Desc        string       `gorm:"column:desc"`
	Users       []User       `gorm:"many2many:role_users;"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
