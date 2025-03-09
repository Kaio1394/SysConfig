package interfaces

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"context"
)

type IAgentRepository interface {
	CreateAgent(ctx context.Context, agent models.Agent) (models.Agent, error)
	GetAgents(ctx context.Context) ([]models.Agent, error)
	UpdateAgent(ctx context.Context, id uint, agentDto dtos.AgentUpdateDto)
}
