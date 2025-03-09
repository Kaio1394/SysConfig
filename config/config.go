package config

type Config struct {
	App App
	Server
	DataBase
	Redis
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
