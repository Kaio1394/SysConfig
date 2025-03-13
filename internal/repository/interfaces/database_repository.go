package interfaces

import (
	"SysConfig/internal/models"
	"context"
)

type DataBaseRepository interface {
	CreateConfigDatabase(ctx context.Context, agent *models.Database) error
	GetConfigDatabase(ctx context.Context) ([]models.Database, error)
	//UpdateConfigDatabase(ctx context.Context, id uint64, agent models.Database) error
	//GetConfigDatabaseById(ctx context.Context, id uint64) (models.Database, error)
}
