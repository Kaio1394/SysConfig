package interfaces

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"context"
)

type DatabaseService interface {
	CreateConfigDatabase(ctx context.Context, database models.Database) error
	GetConfigsDatabase(ctx context.Context) ([]dtos.DatabaseReadDto, error)
	GetConfigDatabaseByUuid(ctx context.Context, uuid string) (dtos.DatabaseReadDto, error)
	UpdateConfigDatabase(ctx context.Context, database models.Database) error
	DeleteConfigDatabase(ctx context.Context, uuid string) error
	GetConfigDatabaseByTag(ctx context.Context, tag string) (dtos.DatabaseReadDto, error)
}
