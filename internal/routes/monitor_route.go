package routes

import (
	"SysConfig/internal/handlers"
	"SysConfig/internal/repository"
	"SysConfig/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterConfigMonitorRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewMonitorRepositoryImpl(db)
	serv := services.NewMonitorServiceImpl(repo)
	h := handlers.NewMonitorHandlerImpl(serv)
	r.POST("/monitor", h.CreateConfigMonitor)
}
