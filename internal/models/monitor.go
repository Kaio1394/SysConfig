package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Monitor struct {
	Uuid      string    `json:"uuid" gorm:"primaryKey"`
	Tag       string    `json:"tag" gorm:"unique;not null"`
	Host      string    `json:"host" gorm:"not null"`
	Port      int       `json:"port" gorm:"not null"`
	EditDate  time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (m *Monitor) BeforeCreate(tx *gorm.DB) (err error) {
	m.Uuid = uuid.New().String()
	return
}

func (Monitor) TableName() string {
	return "t_config_monitors"
}
