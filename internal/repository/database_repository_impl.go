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
func (r *DatabaseRepositoryImpl) UpdateConfigDatabase(ctx context.Context, database *models.Database) error {
	return r.db.WithContext(ctx).Model(&database).Where("uuid = ?", database.Uuid).Updates(database).Error
}

func (r *DatabaseRepositoryImpl) GetConfigDatabaseByUuid(ctx context.Context, uuid string) (models.Database, error) {
	var database models.Database
	if err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&database).Error; err != nil {
		return database, err
	}
	return database, nil
}
