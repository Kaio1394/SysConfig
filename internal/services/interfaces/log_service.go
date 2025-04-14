package interfaces

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"context"
)

type LogService interface {
	GetConfigLogs(ctx context.Context) ([]dtos.LogReadDto, error)
	CreateConfigLog(ctx context.Context, log models.Log) error
	UpdateConfigLog(ctx context.Context, uuid string, logDto dtos.LogUpdateDto) error
	GetConfigLogByUuid(ctx context.Context, uuid string) (models.Log, error)
	GetConfigLogByTag(ctx context.Context, tag string) (dtos.LogReadDto, error)
	DeleteConfigLogByUuid(ctx context.Context, uuid string) error
}
