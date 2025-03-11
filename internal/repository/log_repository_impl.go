package repository

import (
	"SysConfig/internal/models"
	"context"
	"gorm.io/gorm"
)

type LogRepositoryImpl struct {
	db *gorm.DB
}

func NewLogRepositoryImpl(db *gorm.DB) *LogRepositoryImpl {
	return &LogRepositoryImpl{db: db}
}

func (r *LogRepositoryImpl) CreateConfigLog(ctx context.Context, log *models.Log) error {
	return r.db.WithContext(ctx).Create(&log).Error
}

func (r *LogRepositoryImpl) GetConfigLogs(ctx context.Context) ([]models.Log, error) {
	var logs []models.Log
	if err := r.db.WithContext(ctx).Find(&logs).Error; err != nil {
		return logs, err
	}
	return logs, nil
}
