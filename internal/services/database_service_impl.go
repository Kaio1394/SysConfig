package services

import (
	"SysConfig/internal/models"
	"SysConfig/internal/repository"
	"context"
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
