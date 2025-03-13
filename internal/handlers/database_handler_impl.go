package handlers

import (
	"SysConfig/internal/models"
	"SysConfig/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DatabaseHandlerImpl struct {
	s *services.DatabaseServiceImpl
}

func NewDatabaseHandlerImpl(s *services.DatabaseServiceImpl) *DatabaseHandlerImpl {
	return &DatabaseHandlerImpl{s}
}

func (h *DatabaseHandlerImpl) CreateConfigDatabase(c *gin.Context) {
	var databaseConfig models.Database
	if err := c.ShouldBindJSON(&databaseConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.s.CreateConfigDatabase(context.Background(), databaseConfig); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created config database",
	})
}
