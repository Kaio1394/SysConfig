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

type AgentHandlerImpl struct {
	s *services.AgentServiceImpl
}

func NewAgentHandlerImpl(s *services.AgentServiceImpl) *AgentHandlerImpl {
	return &AgentHandlerImpl{s: s}
}

func (h *AgentHandlerImpl) CreateAgent(c *gin.Context) {
	var a models.Agent
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	agent, err := h.s.CreateAgent(context.Background(), a)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"agent": agent,
	})
}

func (h *AgentHandlerImpl) GetAgents(c *gin.Context) {
	agents, err := h.s.GetAgents(context.Background())
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"agents": agents,
	})
}

func (h *AgentHandlerImpl) UpdateAgent(c *gin.Context) {
	var a dtos.AgentUpdateDto
	idStr := c.GetHeader("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := h.s.UpdateAgent(context.Background(), id, a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "agent updated",
	})
}

func (h *AgentHandlerImpl) GetAgentById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	agent, err := h.s.GetAgentById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"agent": agent,
	})
}
