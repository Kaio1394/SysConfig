package interfaces

import (
	"SysConfig/internal/models"
	"context"
)

type DataBaseRepository interface {
	CreateConfigDatabase(ctx context.Context, database *models.Database) error
	GetConfigDatabase(ctx context.Context) ([]models.Database, error)
	UpdateConfigDatabase(ctx context.Context, database *models.Database) error
	GetConfigDatabaseByUuid(ctx context.Context, uuid string) (models.Database, error)
	GetConfigDatabaseByTag(ctx context.Context, tag string) (models.Database, error)
	DeleteConfigDatabase(ctx context.Context, uuid string) error
}
