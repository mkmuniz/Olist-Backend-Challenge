package configs

type Config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string `json:"port"`
}

type DBConfig struct {
	URL  string
	Name string
}
