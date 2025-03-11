package interfaces

import "github.com/gin-gonic/gin"

type AgentHandler interface {
	GetAgentById(c *gin.Context)
	UpdateAgent(c *gin.Context)
	GetAgents(c *gin.Context)
	CreateAgent(c *gin.Context)
}
