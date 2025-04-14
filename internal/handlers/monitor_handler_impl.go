package handlers

import (
	"SysConfig/internal/models"
	"SysConfig/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MonitorHandlerImpl struct {
	s *services.MonitorServiceImpl
}

func NewMonitorHandlerImpl(s *services.MonitorServiceImpl) *MonitorHandlerImpl {
	return &MonitorHandlerImpl{s}
}

func (h *MonitorHandlerImpl) CreateConfigMonitor(c *gin.Context) {
	var monitorModel models.Monitor
	if err := c.ShouldBindJSON(&monitorModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.s.CreateConfigMonitor(context.Background(), &monitorModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Monitor created"})
}
