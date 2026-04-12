package service

import (
	"go-rbac-example/internal/dao"
	"go-rbac-example/internal/model"
)

type DataService struct {
	dao *dao.DataDao
}

func NewDataService(dao *dao.DataDao) *DataService {
	return &DataService{dao: dao}
}

func (s *DataService) ListData(fields []string) ([]model.DataResp, error) {
	return s.dao.ListData(fields)
}
