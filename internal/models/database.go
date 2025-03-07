package models

type Database struct {
	Id               int    `json:"id" gorm:"primaryKey"`
	Tag              string `json:"tag" gorm:"not null"`
	TypeDatabase     string `json:"type_database" gorm:"not null"`
	StringConnection string `json:"string_connection" gorm:"not null"`
}

func (Database) TableName() string {
	return "t_config_databases"
}
