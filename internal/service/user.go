package service

import (
	"go-rbac-example/internal/dao"
	"go-rbac-example/internal/model"
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

func (s *UserService) DeleteUser(ids []uint64) error {
	return s.dao.DeleteUser(ids)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.dao.UpdateUser(user)
}

func (s *UserService) ListUsers(page model.Page) ([]model.User, int64, error) {
	return s.dao.ListUsers(page)
}
