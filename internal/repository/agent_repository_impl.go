package repository

import (
	"SysConfig/internal/models"
	"github.com/jinzhu/copier"
	"time"

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

func (r *AgentRepositoryImpl) UpdateAgent(ctx context.Context, id uint64, agent models.Agent) error {
	var existingAgent models.Agent
	if err := r.db.WithContext(ctx).First(&existingAgent, id).Error; err != nil {
		return err
	}
	_ = copier.CopyWithOption(&existingAgent, &agent, copier.Option{IgnoreEmpty: true})
	agent.EditDate = time.Now()

	if err := r.db.WithContext(ctx).Model(&existingAgent).Where("id = ?", id).Updates(existingAgent).Error; err != nil {
		return err
	}
	return nil
}

func (r *AgentRepositoryImpl) GetAgentById(ctx context.Context, id uint64) (models.Agent, error) {
	var agent models.Agent
	if err := r.db.WithContext(ctx).First(&agent, id).Error; err != nil {
		return models.Agent{}, err
	}
	return agent, nil
}
