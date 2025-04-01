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

func (r *LogRepositoryImpl) UpdateConfigLog(ctx context.Context, log *models.Log) error {
	return r.db.WithContext(ctx).Model(&log).Where("uuid = ?", log.Uuid).Updates(log).Error
}

func (r *LogRepositoryImpl) GetConfigLogByUuid(ctx context.Context, uuid string) (models.Log, error) {
	var log models.Log
	if err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&log).Error; err != nil {
		return log, err
	}
	return log, nil
}

func (r *LogRepositoryImpl) GetConfigLogByTag(ctx context.Context, tag string) (models.Log, error) {
	var log models.Log
	if err := r.db.WithContext(ctx).Where("tag = ?", tag).First(&log).Error; err != nil {
		return log, err
	}
	return log, nil
}
func (r *LogRepositoryImpl) DeleteConfigLogByUuid(ctx context.Context, uuid string) error {
	return r.db.WithContext(ctx).Delete(&models.Log{}, uuid).Error
}
