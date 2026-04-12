package dao

import (
	"go-rbac-example/internal/model"

	"gorm.io/gorm"
)

type DataDao struct {
	db *gorm.DB
}

func NewDataDao(db *gorm.DB) *DataDao {
	return &DataDao{db: db}
}

func (d *DataDao) ListData(fields []string) ([]model.DataResp, error) {
	var data []model.DataResp
	err := d.db.Model(&model.Data{}).Select(fields).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
