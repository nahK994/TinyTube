package app

type AppConfig struct {
	Host    string
	Port    int
	Workers []string
}

func GetConfig() AppConfig {
	return AppConfig{
		Host: "127.0.0.1",
		Port: 8002,
		Workers: []string{
			"http://localhost:8081",
			"http://localhost:8082",
			"http://localhost:8083",
			"http://localhost:8084",
			"http://localhost:8085",
		},
	}
}
