package interfaces

import (
	"SysConfig/internal/models"
	"context"
)

type LogRepository interface {
	CreateConfigLog(ctx context.Context, log *models.Log) error
	GetConfigLogs(ctx context.Context) ([]models.Log, error)
}
