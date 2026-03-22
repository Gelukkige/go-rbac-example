package dao

import (
	"go-rbac-example/internal/model"
	"sort"

	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{db: db}
}

func (dao *RoleDao) CreateRole(req model.RoleCreateReq) error {
	return dao.db.Transaction(func(tx *gorm.DB) error {
		newRole := &model.Role{
			Name: req.Name,
			Desc: req.Desc,
		}

		if err := tx.Create(&newRole).Error; err != nil {
			return err
		}

		if len(req.Permissions) > 0 {
			var permissions []model.Permission
			for _, pReq := range req.Permissions {
				var perm model.Permission
				sort.Strings(pReq.Columns)
				err := tx.Where(model.Permission{
					Page:   pReq.Page,
					Action: pReq.Action,
				}).Attrs(model.Permission{Columns: pReq.Columns}).
					FirstOrCreate(&perm).Error

				if err != nil {
					return err
				}
				permissions = append(permissions, perm)
			}
			if err := tx.Model(&newRole).Association("Permissions").Append(permissions); err != nil {
				return err
			}
		}
		return nil
	})
}

func (dao *RoleDao) DeleteRole(ids []uint64) error {
	return dao.db.Where("id IN ?", ids).Delete(&model.Role{}).Error
}

func (dao *RoleDao) UpdateRole(req *model.RoleUpdateReq) error {
	return dao.db.Transaction(func(tx *gorm.DB) error {
		var role model.Role
		if err := tx.First(&role, req.ID).Error; err != nil {
			return err
		}

		var updateData = make(map[string]interface{})
		if req.Name != "" {
			updateData["name"] = req.Name
		}
		if req.Desc != "" {
			updateData["desc"] = req.Desc
		}

		if len(updateData) > 0 {
			if err := tx.Model(&role).Updates(updateData).Error; err != nil {
				return err
			}
		}

		if req.UserIDs != nil {
			var users []model.User
			if len(req.UserIDs) > 0 {
				if err := tx.Where("id IN ?", req.UserIDs).Find(&users).Error; err != nil {
					return err
				}
			}
			if err := tx.Model(&role).Association("Users").Replace(users); err != nil {
				return err
			}
		}

		if req.Permissions != nil {
			var permissions []model.Permission
			for _, pReq := range *req.Permissions {
				var perm model.Permission
				sort.Strings(pReq.Columns)
				err := tx.Where(model.Permission{
					Page:   pReq.Page,
					Action: pReq.Action,
				}).Attrs(model.Permission{Columns: pReq.Columns}).
					FirstOrCreate(&perm).Error

				if err != nil {
					return err
				}
				permissions = append(permissions, perm)
			}
			if err := tx.Model(&role).Association("Permissions").Replace(permissions); err != nil {
				return err
			}
		}

		return nil
	})
}

func (dao *RoleDao) ListRoles(page model.Page) ([]model.RoleInfoResp, int64, error) {
	var roles []model.RoleInfoResp
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

	err = dao.db.Model(&model.Role{}).Offset((pageNum - 1) * pageSize).Limit(pageSize).Scan(&roles).Error
	if err != nil {
		return nil, 0, err
	}

	for i := range roles {
		if roles[i].Users == nil {
			roles[i].Users = []string{}
		}
		if roles[i].Permissions == nil {
			roles[i].Permissions = []string{}
		}
	}

	return roles, total, nil
}
