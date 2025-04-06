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

func (h *DatabaseHandlerImpl) GetConfigsDatabase(c *gin.Context) {
	list, err := h.s.GetConfigsDatabase(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"configs database": list,
	})
}
func (h *DatabaseHandlerImpl) GetConfigDatabaseByUuid(c *gin.Context) {
	uuid := c.GetHeader("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uuid is required"})
		return
	}
	data, err := h.s.GetConfigDatabaseByUuid(context.Background(), uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"config database": data,
	})
}

func (h *DatabaseHandlerImpl) UpdateConfigDatabase(c *gin.Context) {
	var databaseConfig models.Database
	if err := c.ShouldBindJSON(&databaseConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.s.UpdateConfigDatabase(context.Background(), databaseConfig); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated config database",
	})
}

func (h *DatabaseHandlerImpl) DeleteConfigDatabase(c *gin.Context) {
	uuid := c.GetHeader("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uuid is required"})
		return
	}
	if err := h.s.DeleteConfigDatabase(context.Background(), uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted config database",
	})
}
