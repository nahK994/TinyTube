package app

type AppConfig struct {
	Host                 string
	Port                 int
	JWT_secret_key       []byte
	Bcrypt_password_cost int
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

var userManagementConfig Config = Config{
	App: AppConfig{
		Host:                 "127.0.0.1",
		Port:                 8001,
		JWT_secret_key:       []byte("JWT_secret_key"),
		Bcrypt_password_cost: 14,
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
	return userManagementConfig
}
