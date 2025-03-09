package services

import (
	"SysConfig/internal/models"
	"SysConfig/internal/models/dtos"
	"SysConfig/internal/repository"
	"context"
	"github.com/jinzhu/copier"
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
func (a *AgentServiceImpl) GetAgents(ctx context.Context) ([]dtos.AgentReadDto, error) {
	agents, err := a.r.GetAgents(ctx)
	if err != nil {
		return nil, err
	}
	var listDtos []dtos.AgentReadDto
	_ = copier.CopyWithOption(&listDtos, &agents, copier.Option{IgnoreEmpty: true})
	return listDtos, nil
}

func (a *AgentServiceImpl) UpdateAgent(ctx context.Context, id uint64, agentDto dtos.AgentUpdateDto) error {
	var agent models.Agent
	_ = copier.CopyWithOption(&agent, agentDto, copier.Option{IgnoreEmpty: true})
	err := a.r.UpdateAgent(ctx, id, agent)
	if err != nil {
		return err
	}
	return nil
}

func (a *AgentServiceImpl) GetAgentById(ctx context.Context, id uint64) (dtos.AgentReadDto, error) {
	var agentDto dtos.AgentReadDto
	agentModel, err := a.r.GetAgentById(ctx, id)
	if err != nil {
		return dtos.AgentReadDto{}, err
	}

	_ = copier.CopyWithOption(&agentDto, agentModel, copier.Option{IgnoreEmpty: true})
	return agentDto, nil
}
