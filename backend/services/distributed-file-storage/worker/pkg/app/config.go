package app

type AppConfig struct {
	Host string
	Port int
}

func GetConfig() AppConfig {
	return AppConfig{
		Host: "127.0.0.1",
		Port: 8003,
	}
}
