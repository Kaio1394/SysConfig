package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Agent struct {
	Uuid      string    `json:"uuid" gorm:"primaryKey"`
	Tag       string    `json:"tag" gorm:"unique;not null"`
	Host      string    `json:"host" gorm:"not null"`
	Port      int       `json:"port" gorm:"not null"`
	EditDate  time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (a *Agent) BeforeCreate(tx *gorm.DB) (err error) {
	a.Uuid = uuid.New().String()
	return
}

func (Agent) TableName() string {
	return "t_config_agents"
}
