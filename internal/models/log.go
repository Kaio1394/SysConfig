package models

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
	Id               int    `json:"id" gorm:"primaryKey"`
	Tag              string `json:"tag" gorm:"not null"`
	FileName         string `json:"filename" gorm:"not null"`
	PathDir          string `json:"path_dir" gorm:"not null"`
	Level            int    `json:"level" gorm:"not null"`
	FormatDate       int    `json:"format_date" gorm:"not null"`
	StringConnection string `json:"string_connection" gorm:"not null"`
}

func (Log) TableName() string {
	return "t_config_logs"
}
