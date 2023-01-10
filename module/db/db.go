package db

import (
	"fmt"
	"log"

	"budgetapp/module/consts"
	"budgetapp/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDb(dbConfig config.DatabaseConfig, env consts.Environment) *gorm.DB {
	dbURL := getDbUrl(dbConfig)

	gormConif := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbConfig.Schema + ".",
			SingularTable: true,
		}}

	if env != consts.PRODUCTION {
		gormConif.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dbURL), gormConif)

	if err != nil {
		log.Println("Failed to connect to database", err.Error())
		panic(err)
	}

	fmt.Println("successfully connected to database")

	return db
}
