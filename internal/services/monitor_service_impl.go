package services

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"SysConfig/internal/repository"
	"context"
	"errors"
	"github.com/jinzhu/copier"
)

type MonitorServiceImpl struct {
	r *repository.MonitorRepositoryImpl
}

func NewMonitorServiceImpl(r *repository.MonitorRepositoryImpl) *MonitorServiceImpl {
	return &MonitorServiceImpl{r: r}
}

func (s *MonitorServiceImpl) GetConfigMonitorByTag(ctx context.Context, tag string) (dtos.MonitorReadDto, error) {
	var monitorDto dtos.MonitorReadDto
	monitorModel, err := s.r.GetConfigMonitorByTag(ctx, tag)
	if err != nil {
		return monitorDto, err
	}
	_ = copier.CopyWithOption(&monitorDto, &monitorModel, copier.Option{IgnoreEmpty: true})
	return monitorDto, nil
}

func (s *MonitorServiceImpl) CreateConfigMonitor(ctx context.Context, monitor *models.Monitor) error {
	_, err := s.GetConfigMonitorByTag(ctx, monitor.Tag)
	if err == nil {
		return errors.New("monitor tag already exists")
	}
	err = s.r.CreateConfigMonitor(ctx, monitor)
	if err != nil {
		return err
	}
	return nil
}

func (s *MonitorServiceImpl) UpdateConfigMonitor(ctx context.Context, monitor *models.Monitor) error {
	_, err := s.GetConfigMonitorByTag(ctx, monitor.Tag)
	if err != nil {
		return errors.New("monitor tag does not exist")
	}
	err = s.r.UpdateConfigMonitor(ctx, monitor)
	if err != nil {
		return err
	}
	return nil
}
