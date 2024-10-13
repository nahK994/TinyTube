package app

type AppConfig struct {
	Host                   string
	Port                   int
	JWT_secret_key         []byte
	JWT_exp_minutes        int
	RefreshToken_exp_hours int
	Bcrypt_password_cost   int
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
		Host:                   "127.0.0.1",
		Port:                   8000,
		JWT_secret_key:         []byte("JWT_secret_key"),
		JWT_exp_minutes:        15,
		RefreshToken_exp_hours: 7 * 24,
		Bcrypt_password_cost:   14,
	},
	Database: DBConfig{
		Username: "user",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     5000,
		Name:     "auth_db",
	},
}

func GetConfig() Config {
	return appConfig
}
