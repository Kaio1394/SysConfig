package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

/*
PanicLevel = 0
FatalLevel = 1
ErrorLevel = 2
WarnLevel = 3
InfoLevel = 4
DebugLevel = 5
TraceLevel = 6
*/

type Log struct {
	Uuid       string    `json:"uuid" gorm:"primaryKey"`
	Tag        string    `json:"tag" gorm:"unique;not null"`
	FileName   string    `json:"filename" gorm:"unique;not null"`
	PathDir    string    `json:"path_dir" gorm:"unique;not null"`
	Level      int       `json:"level" gorm:"not null"`
	FormatDate string    `json:"format_date" gorm:"not null"`
	EditDate   time.Time `gorm:"autoUpdateTime"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (l *Log) BeforeCreate(tx *gorm.DB) (err error) {
	l.Uuid = uuid.New().String()
	return
}

func (Log) TableName() string {
	return "t_config_logs"
}
