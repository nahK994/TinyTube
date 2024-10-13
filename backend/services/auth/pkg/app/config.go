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
		Port: 8000,
	},
	Database: DBConfig{
		Username: "user",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     5432,
		Name:     "auth_db",
	},
}

func GetConfig() Config {
	return appConfig
}
