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

func (dao *UserDao) DeleteUser(ids []uint64) error {
	return dao.db.Where("id IN ?", ids).Delete(&model.User{}).Error
}

func (dao *UserDao) UpdateUser(user *model.User) error {
	return dao.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (dao *UserDao) ListUsers(page model.Page) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	err := dao.db.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	pageNum := page.PageNum
	pageSize := page.PageSize

	if pageNum < 0 {
		pageNum = model.DefaultPageNum
	}
	if pageSize <= 0 {
		pageSize = model.DefaultPageSize
	}
	err = dao.db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
