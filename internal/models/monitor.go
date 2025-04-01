package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Monitor struct {
	Uuid     string    `json:"uuid" gorm:"primaryKey"`
	Host     string    `json:"host" gorm:"not null"`
	Port     int       `json:"port" gorm:"not null"`
	EditDate time.Time `gorm:"autoUpdateTime"`
}

func (m *Monitor) BeforeCreate(tx *gorm.DB) (err error) {
	m.Uuid = uuid.New().String()
	return
}

func (Agent) Monitor() string {
	return "t_config_monitor"
}
