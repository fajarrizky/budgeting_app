package config


type DatabaseConfig struct {
	Host     string
	DbName   string
	Port     int
	User     string
	Password string
	Schema   string
}

func newDatabaseConfig(cm ConfigManager) DatabaseConfig {
	return DatabaseConfig{
		Host:     cm.ReadEnvString("DB_HOST"),
		DbName:   cm.ReadEnvString("DB_NAME"),
		Port:     cm.ReadEnvInt("DB_PORT"),
		User:     cm.ReadEnvString("DB_USER"),
		Password: cm.ReadEnvString("DB_PASSWORD"),
		Schema:   cm.ReadEnvString("DB_SCHEMA"),
	}
}
