package config

// Config is a struct
type Config struct {
	DB *DBConfig
}

// DBConfig is a struct
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

// GetConfig is a config settings
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "secret",
			Name:     "employee_db",
			Charset:  "utf8",
		},
	}
}
