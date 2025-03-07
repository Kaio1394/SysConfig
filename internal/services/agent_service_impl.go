package services

import (
	"SysConfig/internal/models"
	"SysConfig/internal/repository"
	"context"
)

type AgentServiceImpl struct {
	r *repository.AgentRepositoryImpl
}

func NewAgentServiceImpl(r *repository.AgentRepositoryImpl) *AgentServiceImpl {
	return &AgentServiceImpl{r: r}
}

func (a *AgentServiceImpl) CreateAgent(ctx context.Context, agent models.Agent) (models.Agent, error) {
	return a.r.CreateAgent(ctx, agent)
}
func (a *AgentServiceImpl) GetAgents(ctx context.Context) ([]models.Agent, error) {
	return a.r.GetAgents(ctx)
}
