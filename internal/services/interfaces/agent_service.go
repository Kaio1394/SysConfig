package interfaces

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"context"
)

type AgentService interface {
	UpdateAgent(ctx context.Context, id uint64, agentDto dtos.AgentUpdateDto) error
	CreateAgent(ctx context.Context, agent models.Agent) (models.Agent, error)
	GetAgents(ctx context.Context) ([]dtos.AgentReadDto, error)
}
