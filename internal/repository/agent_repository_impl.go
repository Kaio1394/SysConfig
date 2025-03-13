package repository

import (
	"SysConfig/internal/models"
	"context"
	"gorm.io/gorm"
)

type AgentRepositoryImpl struct {
	db *gorm.DB
}

func NewAgentRepositoryImpl(db *gorm.DB) *AgentRepositoryImpl {
	return &AgentRepositoryImpl{db: db}
}

func (r *AgentRepositoryImpl) CreateAgent(ctx context.Context, agent models.Agent) (models.Agent, error) {
	if err := r.db.WithContext(ctx).Create(&agent).Error; err != nil {
		return agent, err
	}
	return agent, nil
}

func (r *AgentRepositoryImpl) GetAgents(ctx context.Context) ([]models.Agent, error) {
	var agents []models.Agent
	if err := r.db.WithContext(ctx).Find(&agents).Error; err != nil {
		return agents, err
	}
	return agents, nil
}

func (r *AgentRepositoryImpl) UpdateAgent(ctx context.Context, agent models.Agent) error {
	return r.db.WithContext(ctx).Model(&agent).Where("id = ?", agent.Id).Updates(agent).Error
}

func (r *AgentRepositoryImpl) GetAgentById(ctx context.Context, id uint64) (models.Agent, error) {
	var agent models.Agent
	if err := r.db.WithContext(ctx).First(&agent, id).Error; err != nil {
		return models.Agent{}, err
	}
	return agent, nil
}
