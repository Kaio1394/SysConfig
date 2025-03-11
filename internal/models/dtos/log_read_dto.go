package dtos

type LogReadDto struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Tag        string `json:"tag" gorm:"not null"`
	FileName   string `json:"filename" gorm:"not null"`
	PathDir    string `json:"path_dir" gorm:"not null"`
	Level      int    `json:"level" gorm:"not null"`
	FormatDate string `json:"format_date" gorm:"not null"`
}
