package config

type Config struct {
	App App
	Server
	DataBase
	Redis
	ConfigAuth
}
type App struct {
	Name string
}

type Server struct {
	Host string
	Port int
}

type DataBase struct {
	TypeDatabase     string
	StringConnection string
}

type Redis struct {
	Host string
	Port int
}

type ConfigAuth struct {
	Host     string
	Port     int
	Endpoint string
}
