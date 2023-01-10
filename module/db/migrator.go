package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"budgetapp/config"
	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"
)

func RunMigrations(dbConfig config.DatabaseConfig, dir string) {
	dbUrl := getDbUrl(dbConfig)

	dbURL := fmt.Sprintf("%s search_path=%s", dbUrl, dbConfig.Schema)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("failed to connect to db for migration | error: %v", err.Error())
		panic(err)
	}

	defer db.Close()
	goose.SetLogger(log.New(os.Stdout, "", log.LstdFlags))

	if err := goose.SetDialect("postgres"); err != nil {
		log.Printf("failed set dialect for migration")
		panic(err)
	}

	if err := goose.Up(db, dir, goose.WithAllowMissing()); err != nil {
		log.Printf("failed to run migration| error: %v", err.Error())
		panic(err)
	}

	fmt.Println("db migration completed!")
}
