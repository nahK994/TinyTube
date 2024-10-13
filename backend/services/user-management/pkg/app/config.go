package app

type AppConfig struct {
	Host string
	Port int
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}

type Config struct {
	App      AppConfig
	Database DBConfig
}

var appConfig Config = Config{
	App: AppConfig{
		Host: "127.0.0.1",
		Port: 8001,
	},
	Database: DBConfig{
		Username: "user",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     5001,
		Name:     "user_management_db",
	},
}

func GetConfig() Config {
	return appConfig
}
