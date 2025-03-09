package routes

import (
	"SysConfig/internal/handlers"
	"SysConfig/internal/repository"
	"SysConfig/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouteAgent(r *gin.Engine, db *gorm.DB) {
	rep := repository.NewAgentRepositoryImpl(db)
	s := services.NewAgentServiceImpl(rep)
	h := handlers.NewAgentHandlerImpl(s)
	r.GET("/agents", h.GetAgents)
	r.POST("/agent", h.CreateAgent)
	r.PUT("/agent", h.UpdateAgent)
}
