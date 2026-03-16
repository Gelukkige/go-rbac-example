package dao

import (
	"go-rbac-example/internal/model"

	"gorm.io/gorm"
)

type PermissionDao struct {
	db *gorm.DB
}

func NewPermissionDao(db *gorm.DB) *PermissionDao {
	return &PermissionDao{db: db}
}

func (dao *PermissionDao) CreatePermission(permission *model.Permission) error {
	return dao.db.Create(permission).Error
}

func (dao *PermissionDao) DeletePermission(ids []uint64) error {
	return dao.db.Where("id IN ?", ids).Delete(&model.Permission{}).Error
}

func (dao *PermissionDao) UpdatePermission(permission *model.Permission) error {
	return dao.db.Model(&model.Permission{}).Where("id = ?", permission.ID).Updates(permission).Error
}

func (dao *PermissionDao) ListPermissions(page model.Page) ([]model.Permission, int64, error) {
	var permissions []model.Permission
	var total int64

	err := dao.db.Model(&model.Permission{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	pageNum := page.PageNum
	pageSize := page.PageSize

	if pageNum < 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	err = dao.db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&permissions).Error
	if err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}
