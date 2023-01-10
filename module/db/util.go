package db

import (
	"fmt"

	"budgetapp/config"
)

func getDbUrl(dbConfig config.DatabaseConfig) string {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DbName,
		dbConfig.Password,
	)

	return dbURL
}
