package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Database struct {
	Uuid             string `json:"uuid" gorm:"primaryKey"`
	Tag              string `json:"tag" gorm:"not null"`
	TypeDatabase     string `json:"type_database" gorm:"not null"`
	StringConnection string `json:"string_connection" gorm:"not null"`
}

func (d *Database) BeforeCreate(tx *gorm.DB) (err error) {
	d.Uuid = uuid.New().String()
	return
}

func (Database) TableName() string {
	return "t_config_databases"
}
