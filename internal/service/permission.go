package service

// import (
// 	"go-rbac-example/internal/dao"
// 	"go-rbac-example/internal/model"
// )

// type PermissionService struct {
// 	dao *dao.PermissionDao
// }

// func NewPermissionService(dao *dao.PermissionDao) *PermissionService {
// 	return &PermissionService{dao: dao}
// }

// func (s *PermissionService) CreatePermission(req model.PermissionCreateReq) error {
// 	Permission := model.Permission{
// 		Page:    req.Page,
// 		Action:  req.Action,
// 		Columns: req.Columns,
// 	}
// 	return s.dao.CreatePermission(&Permission)
// }

// func (s *PermissionService) DeletePermission(req model.DeleteIDs) error {
// 	return s.dao.DeletePermission(req.IDs)
// }

// func (s *PermissionService) UpdatePermission(req model.PermissionUpdateReq) error {
// 	Permission := model.Permission{
// 		ID:      req.ID,
// 		Page:    req.Page,
// 		Action:  req.Action,
// 		Columns: req.Columns,
// 	}
// 	return s.dao.UpdatePermission(&Permission)
// }

// func (s *PermissionService) ListPermissions(page model.Page) ([]model.Permission, int64, error) {
// 	return s.dao.ListPermissions(page)
// }
