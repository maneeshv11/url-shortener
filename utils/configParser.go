package utils

type AppConfig struct {
	Port     string
	RedisUri string
}

var appConfig AppConfig

func LoadConfig() AppConfig {
	appConfig = AppConfig{
		Port:     "8080",
		RedisUri: "127.0.0.1:6379",
	}

	return appConfig
}

func GetConfig() AppConfig {
	return appConfig
}
