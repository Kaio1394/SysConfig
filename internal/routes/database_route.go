package routes

import (
	"SysConfig/internal/handlers"
	"SysConfig/internal/repository"
	"SysConfig/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterDatabaseRoute(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewDatabaseRepositoryImpl(db)
	s := services.NewDatabaseServiceImpl(repo)
	h := handlers.NewDatabaseHandlerImpl(s)

	r.POST("/database/config", h.CreateConfigDatabase)
	r.GET("/database/config", h.GetConfigsDatabase)
}
