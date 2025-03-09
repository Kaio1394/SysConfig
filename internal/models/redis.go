package models

import "time"

type Redis struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	Host     string    `json:"host" gorm:"not null"`
	Port     int       `json:"port" gorm:"not null"`
	EditDate time.Time `gorm:"autoUpdateTime"`
}

func (Redis) TableName() string {
	return "t_config_redis"
}
