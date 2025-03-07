package models

import "time"

type Monitor struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	Host     string    `json:"host" gorm:"not null"`
	Port     int       `json:"port" gorm:"not null"`
	EditDate time.Time `gorm:"autoUpdateTime"`
}

func (Agent) Monitor() string {
	return "t_config_monitor"
}
