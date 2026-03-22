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
	return s.dao.CreateRole(req)
}

func (s *RoleService) DeleteRole(req model.DeleteIDs) error {
	return s.dao.DeleteRole(req.IDs)
}

func (s *RoleService) UpdateRole(req model.RoleUpdateReq) error {
	return s.dao.UpdateRole(&req)
}

func (s *RoleService) ListRoles(page model.Page) ([]model.RoleInfoResp, int64, error) {
	return s.dao.ListRoles(page)
}
