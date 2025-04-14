package repository

import (
	"SysConfig/internal/models"
	"context"
	"gorm.io/gorm"
)

type MonitorRepositoryImpl struct {
	db *gorm.DB
}

func NewMonitorRepositoryImpl(db *gorm.DB) *MonitorRepositoryImpl {
	return &MonitorRepositoryImpl{db: db}
}

func (r *MonitorRepositoryImpl) CreateConfigMonitor(ctx context.Context, monitor *models.Monitor) error {
	return r.db.WithContext(ctx).Create(monitor).Error
}
func (r *MonitorRepositoryImpl) GetAllConfigMonitor(ctx context.Context) ([]models.Monitor, error) {
	var monitors []models.Monitor
	if err := r.db.WithContext(ctx).Find(&monitors).Error; err != nil {
		return monitors, err
	}
	return monitors, nil
}

func (r *MonitorRepositoryImpl) UpdateConfigMonitor(ctx context.Context, monitor *models.Monitor) error {
	return r.db.WithContext(ctx).Model(monitor).Updates(monitor).Error
}

func (r *MonitorRepositoryImpl) GetConfigMonitorByTag(ctx context.Context, tag string) (models.Monitor, error) {
	var monitor models.Monitor
	if err := r.db.WithContext(ctx).Where("tag = ?", tag).First(&monitor).Error; err != nil {
		return monitor, err
	}
	return monitor, nil
}
