package models

import "time"

type Agent struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Host      string    `json:"host" gorm:"not null"`
	Port      int       `json:"port" gorm:"not null"`
	EditDate  time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (Agent) TableName() string {
	return "t_config_agents"
}
