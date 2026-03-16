package service

import (
	"go-rbac-example/internal/dao"
	"go-rbac-example/internal/model"
)

type RoleService struct {
	dao *dao.RoleDao
}

func NewRoleService(dao *dao.RoleDao) *RoleService {
	return &RoleService{dao: dao}
}

func (s *RoleService) CreateRole(req model.RoleCreateReq) error {
	users := make([]model.User, len(req.UserIDs))
	for i, userID := range req.UserIDs {
		users[i] = model.User{ID: userID}
	}
	permissions := make([]model.Permission, len(req.PermissionIDs))
	for i, permID := range req.PermissionIDs {
		permissions[i] = model.Permission{ID: permID}
	}
	role := model.Role{
		Name:        req.Name,
		Desc:        req.Desc,
		Users:       users,
		Permissions: permissions,
	}
	return s.dao.CreateRole(&role)
}

func (s *RoleService) DeleteRole(req model.DeleteIDs) error {
	return s.dao.DeleteRole(req.IDs)
}

func (s *RoleService) UpdateRole(req model.RoleUpdateReq) error {
	users := make([]model.User, len(req.UserIDs))
	for i, userID := range req.UserIDs {
		users[i] = model.User{ID: userID}
	}
	permissions := make([]model.Permission, len(req.PermissionIDs))
	for i, permID := range req.PermissionIDs {
		permissions[i] = model.Permission{ID: permID}
	}
	role := model.Role{
		ID:          req.ID,
		Name:        req.Name,
		Desc:        req.Desc,
		Users:       users,
		Permissions: permissions,
	}
	return s.dao.UpdateRole(&role)
}

func (s *RoleService) ListRoles(page model.Page) ([]model.Role, int64, error) {
	return s.dao.ListRoles(page)
}
