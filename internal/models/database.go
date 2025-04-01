package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	Uuid             string    `json:"uuid" gorm:"primaryKey"`
	Tag              string    `json:"tag" gorm:"unique;not null"`
	TypeDatabase     string    `json:"type_database" gorm:"not null"`
	StringConnection string    `json:"string_connection" gorm:"not null"`
	EditDate         time.Time `gorm:"autoUpdateTime"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
}

func (d *Database) BeforeCreate(tx *gorm.DB) (err error) {
	d.Uuid = uuid.New().String()
	return
}

func (Database) TableName() string {
	return "t_config_databases"
}
