package dtos

type AgentUpdateDto struct {
	Tag  string `json:"tag"`
	Host string `json:"host"`
	Port int    `json:"port"`
}
