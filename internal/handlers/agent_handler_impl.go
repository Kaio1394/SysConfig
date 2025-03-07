package handlers

import (
	"SysConfig/internal/models"
	"SysConfig/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
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
