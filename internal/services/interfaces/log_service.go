package interfaces

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"context"
)

type LogService interface {
	GetConfigLogs(ctx context.Context) ([]dtos.LogReadDto, error)
	CreateConfigLog(ctx context.Context, log models.Log) error
}
