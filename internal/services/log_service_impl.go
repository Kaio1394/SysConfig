package services

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"SysConfig/internal/repository"
	"context"
	"errors"
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

func (s *LogServiceImpl) UpdateConfigLog(ctx context.Context, uuid string, logDto dtos.LogUpdateDto) error {
	logModel, err := s.repo.GetConfigLogByUuid(ctx, uuid)
	if err != nil {
		return err
	}
	_ = copier.CopyWithOption(&logModel, logDto, copier.Option{IgnoreEmpty: true})

	return s.repo.UpdateConfigLog(ctx, &logModel)
}

func (s *LogServiceImpl) GetConfigLogByUuid(ctx context.Context, uuid string) (models.Log, error) {
	return s.repo.GetConfigLogByUuid(ctx, uuid)
}

func (s *LogServiceImpl) GetConfigLogByTag(ctx context.Context, tag string) (dtos.LogReadDto, error) {
	var dtoModel dtos.LogReadDto
	if tag == "" {
		return dtos.LogReadDto{}, errors.New("tag can't be empty")
	}
	model, err := s.repo.GetConfigLogByTag(ctx, tag)
	if err != nil {
		return dtos.LogReadDto{}, err
	}
	_ = copier.CopyWithOption(&dtoModel, &model, copier.Option{IgnoreEmpty: true})
	return dtoModel, nil
}

func (s *LogServiceImpl) DeleteConfigLogByUuid(ctx context.Context, uuid string) error {
	_, err := s.repo.GetConfigLogByUuid(ctx, uuid)
	if err != nil {
		return err
	}
	return s.repo.DeleteConfigLogByUuid(ctx, uuid)
}
