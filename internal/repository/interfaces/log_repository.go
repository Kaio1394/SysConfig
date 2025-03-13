package interfaces

import (
	"SysConfig/internal/models"
	"context"
)

type LogRepository interface {
	CreateConfigLog(ctx context.Context, log *models.Log) error
	GetConfigLogs(ctx context.Context) ([]models.Log, error)
	GetConfigLogById(ctx context.Context, id int) (models.Log, error)
	UpdateConfigLog(ctx context.Context, log *models.Log) error
}
