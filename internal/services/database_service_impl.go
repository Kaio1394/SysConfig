package services

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"SysConfig/internal/repository"
	"context"
	"errors"
	"github.com/jinzhu/copier"
)

type DatabaseServiceImpl struct {
	r *repository.DatabaseRepositoryImpl
}

func NewDatabaseServiceImpl(r *repository.DatabaseRepositoryImpl) *DatabaseServiceImpl {
	return &DatabaseServiceImpl{r: r}
}

func (s *DatabaseServiceImpl) CreateConfigDatabase(ctx context.Context, database models.Database) error {
	return s.r.CreateConfigDatabase(ctx, &database)
}

func (s *DatabaseServiceImpl) GetConfigsDatabase(ctx context.Context) ([]dtos.DatabaseReadDto, error) {
	var configsDto []dtos.DatabaseReadDto
	configsModel, err := s.r.GetConfigDatabase(ctx)
	if err != nil {
		return []dtos.DatabaseReadDto{}, err
	}
	_ = copier.CopyWithOption(&configsDto, &configsModel, copier.Option{DeepCopy: true})
	return configsDto, nil
}

func (s *DatabaseServiceImpl) GetConfigDatabaseByUuid(ctx context.Context, uuid string) (dtos.DatabaseReadDto, error) {
	var configDto dtos.DatabaseReadDto
	configModel, err := s.r.GetConfigDatabaseByUuid(ctx, uuid)
	if err != nil {
		return dtos.DatabaseReadDto{}, err
	}
	_ = copier.Copy(&configDto, &configModel)
	return configDto, nil
}

func (s *DatabaseServiceImpl) UpdateConfigDatabase(ctx context.Context, database models.Database) error {
	_, err := s.GetConfigDatabaseByUuid(ctx, database.Uuid)
	if err != nil {
		return errors.New("database not found")
	}
	return s.r.UpdateConfigDatabase(ctx, &database)
}

func (s *DatabaseServiceImpl) DeleteConfigDatabase(ctx context.Context, uuid string) error {
	_, err := s.GetConfigDatabaseByUuid(ctx, uuid)
	if err != nil {
		return errors.New("database not found")
	}
	return s.r.DeleteConfigDatabase(ctx, uuid)
}

func (s *DatabaseServiceImpl) GetConfigDatabaseByTag(ctx context.Context, tag string) (dtos.DatabaseReadDto, error) {
	var configDto dtos.DatabaseReadDto
	configDatabase, err := s.r.GetConfigDatabaseByTag(ctx, tag)
	if err != nil {
		return configDto, err
	}
	_ = copier.Copy(&configDto, &configDatabase)
	return configDto, nil
}
