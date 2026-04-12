package service

import (
	"context"
	"go-rbac-example/internal/dao"
	"go-rbac-example/internal/model"
	"go-rbac-example/internal/permission"
)

type UserService struct {
	dao *dao.UserDao
}

func NewUserService(dao *dao.UserDao) *UserService {
	return &UserService{dao: dao}
}

func (s *UserService) CreateUser(req model.UserCreateReq) error {
	roles := make([]model.Role, len(req.RoleIDs))
	for i, roleID := range req.RoleIDs {
		roles[i] = model.Role{ID: roleID}
	}
	user := model.User{
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
		Roles: roles,
	}
	return s.dao.CreateUser(&user)
}

func (s *UserService) DeleteUser(req model.DeleteIDs) error {
	err := s.dao.DeleteUser(req.IDs)
	if err != nil {
		return err
	}

	// 删除相关权限缓存
	ctx := context.Background()
	for _, id := range req.IDs {
		permission.DeleteUserCache(ctx, id)
	}
	return nil
}

func (s *UserService) UpdateUser(req model.UserUpdateReq) error {
	roles := make([]model.Role, len(req.RoleIDs))
	for i, roleID := range req.RoleIDs {
		roles[i] = model.Role{ID: roleID}
	}
	user := model.User{
		ID:    req.ID,
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
		Roles: roles,
	}
	err := s.dao.UpdateUser(&user)
	if err != nil {
		return err
	}

	// 删除相关权限缓存
	ctx := context.Background()
	permission.DeleteUserCache(ctx, user.ID)
	return nil
}

func (s *UserService) ListUsers(page model.Page) ([]model.User, int64, error) {
	return s.dao.ListUsers(page)
}
