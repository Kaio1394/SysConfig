package dtos

type AgentReadDto struct {
	Id   int    `json:"id"`
	Tag  string `json:"tag"`
	Host string `json:"host"`
	Port int    `json:"port"`
}
