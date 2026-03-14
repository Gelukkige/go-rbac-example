package dao

import (
	"go-rbac-example/internal/model"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.db.Create(user).Error
}
