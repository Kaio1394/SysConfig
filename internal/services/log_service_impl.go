package services

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"SysConfig/internal/repository"
	"context"
	"github.com/jinzhu/copier"
)

type LogServiceImpl struct {
	repo *repository.LogRepositoryImpl
}

func NewLogServiceImpl(repo *repository.LogRepositoryImpl) *LogServiceImpl {
	return &LogServiceImpl{repo: repo}
}

func (s *LogServiceImpl) GetConfigLogs(ctx context.Context) ([]dtos.LogReadDto, error) {
	var listDtoConfigLog []dtos.LogReadDto
	list, err := s.repo.GetConfigLogs(ctx)
	if err != nil {
		return listDtoConfigLog, err
	}
	_ = copier.CopyWithOption(&listDtoConfigLog, &list, copier.Option{IgnoreEmpty: true})
	return listDtoConfigLog, nil
}

func (s *LogServiceImpl) CreateConfigLog(ctx context.Context, log models.Log) error {
	return s.repo.CreateConfigLog(ctx, &log)
}
