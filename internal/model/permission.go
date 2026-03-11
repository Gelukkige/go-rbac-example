package model

type Permission struct {
	ID      uint64   `gorm:"primaryKey"`
	Page    string   `gorm:"column:page;not null"`
	Action  string   `gorm:"column:action;not null"`
	Columns []string `gorm:"column:columns;type:jsonb;serializer:json"`
	Roles   []Role   `gorm:"many2many:role_permissions;"`
}
