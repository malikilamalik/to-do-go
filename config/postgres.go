package config

import (
	"fmt"
	"log"
	"os"

	"xorm.io/xorm"
)

type Postgres struct {
	*xorm.Engine
}

func dbConfig() string {
	return fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
}

func InitDatabase() (db *Postgres, err error) {
	dbconn, _ := xorm.NewEngine("postgres", dbConfig())
	if err != nil {
		log.Panicf("Failed to connect to database")
	}
	return &Postgres{dbconn}, nil
}
