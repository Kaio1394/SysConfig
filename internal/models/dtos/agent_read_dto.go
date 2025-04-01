package dtos

type AgentReadDto struct {
	Uuid string `json:"uuid"`
	Tag  string `json:"tag"`
	Host string `json:"host"`
	Port int    `json:"port"`
}
