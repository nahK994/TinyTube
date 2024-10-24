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

type MQConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

type Config struct {
	App AppConfig
	DB  DBConfig
	MQ  MQConfig
}

var userManagementConfig Config = Config{
	App: AppConfig{
		Host:                 "127.0.0.1",
		Port:                 8001,
		JWT_secret_key:       []byte("JWT_secret_key"),
		Bcrypt_password_cost: 14,
	},
	DB: DBConfig{
		Username: "user",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     5001,
		Name:     "user_management_db",
	},
	MQ: MQConfig{
		Username: "guest",
		Password: "guest",
		Host:     "127.0.0.1",
		Port:     5672,
	},
}

func GetConfig() Config {
	return userManagementConfig
}
