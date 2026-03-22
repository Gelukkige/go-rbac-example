package dao

import (
	"go-rbac-example/internal/model"

	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{db: db}
}

func (dao *RoleDao) CreateRole(role *model.Role) error {
	return dao.db.Create(role).Error
}

func (dao *RoleDao) DeleteRole(ids []uint64) error {
	return dao.db.Where("id IN ?", ids).Delete(&model.Role{}).Error
}

func (dao *RoleDao) UpdateRole(role *model.Role) error {
	return dao.db.Model(&model.Role{}).Where("id = ?", role.ID).Updates(role).Error
}

func (dao *RoleDao) ListRoles(page model.Page) ([]model.Role, int64, error) {
	var roles []model.Role
	var total int64

	err := dao.db.Model(&model.Role{}).Count(&total).Error
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
	err = dao.db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&roles).Error
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}
