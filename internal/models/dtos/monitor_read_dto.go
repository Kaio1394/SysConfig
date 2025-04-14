package dtos

type MonitorReadDto struct {
	Uuid string `json:"uuid" gorm:"primaryKey"`
	Tag  string `json:"tag" gorm:"unique;not null"`
	Host string `json:"host" gorm:"not null"`
	Port int    `json:"port" gorm:"not null"`
}
