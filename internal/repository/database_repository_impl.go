package repository

import (
	"SysConfig/internal/models"
	"context"
	"gorm.io/gorm"
)

type DatabaseRepositoryImpl struct {
	db *gorm.DB
}

func NewDatabaseRepositoryImpl(db *gorm.DB) *DatabaseRepositoryImpl {
	return &DatabaseRepositoryImpl{db: db}
}

func (r *DatabaseRepositoryImpl) CreateConfigDatabase(ctx context.Context, database *models.Database) error {
	return r.db.WithContext(ctx).Create(&database).Error
}

func (r *DatabaseRepositoryImpl) GetConfigDatabase(ctx context.Context) ([]models.Database, error) {
	var databaseList []models.Database
	if err := r.db.WithContext(ctx).Find(&databaseList).Error; err != nil {
		return databaseList, err
	}
	return databaseList, nil
}
