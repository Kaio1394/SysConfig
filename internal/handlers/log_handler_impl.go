package handlers

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"SysConfig/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LogHandlerImpl struct {
	s *services.LogServiceImpl
}

func NewLogHandlerImpl(s *services.LogServiceImpl) *LogHandlerImpl {
	return &LogHandlerImpl{s}
}

func (h *LogHandlerImpl) CreateConfigLog(c *gin.Context) {
	var configLog models.Log
	if err := c.ShouldBindJSON(&configLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.s.CreateConfigLog(context.Background(), configLog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func (h *LogHandlerImpl) GetConfigLog(c *gin.Context) {
	list, err := h.s.GetConfigLogs(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"config logs": list,
	})
}

func (h *LogHandlerImpl) UpdateConfigLog(c *gin.Context) {
	var configLog dtos.LogUpdateDto
	idStr := c.GetHeader("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&configLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.s.UpdateConfigLog(context.Background(), id, configLog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *LogHandlerImpl) GetConfigLogById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log, err := h.s.GetConfigLogById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"log": log,
	})
}

func (h *LogHandlerImpl) GetConfigLogByTag(c *gin.Context) {
	tag := c.GetHeader("tag")
	log, err := h.s.GetConfigLogByTag(context.Background(), tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"log": log,
	})
}

func (h *LogHandlerImpl) DeleteConfigLogById(c *gin.Context) {
	idStr := c.GetHeader("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.s.DeleteConfigLogById(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
