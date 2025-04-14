package interfaces

import (
	"SysConfig/internal/models"
	"context"
)

type LogRepository interface {
	CreateConfigLog(ctx context.Context, log *models.Log) error
	GetConfigLogs(ctx context.Context) ([]models.Log, error)
	UpdateConfigLog(ctx context.Context, log *models.Log) error
	GetConfigLogByUuid(ctx context.Context, uuid string) (models.Log, error)
	GetConfigLogByTag(ctx context.Context, tag string) (models.Log, error)
	DeleteConfigLogByUuid(ctx context.Context, uuid string) error
}
