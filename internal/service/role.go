package service

import (
	"context"
	"go-rbac-example/internal/dao"
	"go-rbac-example/internal/model"
	"go-rbac-example/internal/permission"
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
	err := s.dao.DeleteRole(req.IDs)
	if err != nil {
		return err
	}

	// 删除相关权限缓存
	ctx := context.Background()
	for _, id := range req.IDs {
		permission.DeleteRoleCache(ctx, id)
	}
	return nil
}
func (s *RoleService) UpdateRole(req model.RoleUpdateReq) error {
	err := s.dao.UpdateRole(&req)
	if err != nil {
		return err
	}

	// 删除相关权限缓存
	ctx := context.Background()
	permission.DeleteRoleCache(ctx, req.ID)
	return nil
}

func (s *RoleService) ListRoles(page model.Page) ([]model.RoleInfoResp, int64, error) {
	return s.dao.ListRoles(page)
}
