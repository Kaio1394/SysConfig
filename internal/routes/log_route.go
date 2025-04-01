package routes

import (
	"SysConfig/internal/handlers"
	"SysConfig/internal/repository"
	"SysConfig/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouteLog(r *gin.Engine, db *gorm.DB) {
	rep := repository.NewLogRepositoryImpl(db)
	s := services.NewLogServiceImpl(rep)
	h := handlers.NewLogHandlerImpl(s)
	r.GET("/logs/config", h.GetConfigLog)
	r.GET("/logs/config/:id", h.GetConfigLogByUuid)
	r.GET("/logs/config/tag", h.GetConfigLogByTag)

	r.POST("/logs/config", h.CreateConfigLog)
	r.DELETE("/logs/config", h.DeleteConfigLogByUuid)
	r.PUT("/logs/config", h.UpdateConfigLog)
}
