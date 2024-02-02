package config

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	SECRET_KEY  string
}

func InitConfig() Config {
	return Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_HOST:     "localhost",
		DB_PORT:     "3306",
		DB_NAME:     "backendtest",
		SECRET_KEY:  "secret",
	}
}
