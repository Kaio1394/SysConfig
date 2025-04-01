package services

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"SysConfig/internal/repository"
	"context"
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
