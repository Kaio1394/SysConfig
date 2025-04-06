package dtos

type DatabaseUpdateDto struct {
	Uuid             string `json:"uuid" gorm:"primaryKey"`
	Tag              string `json:"tag" gorm:"unique;not null"`
	TypeDatabase     string `json:"type_database" gorm:"not null"`
	StringConnection string `json:"string_connection" gorm:"not null"`
}
