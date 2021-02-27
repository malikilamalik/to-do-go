package config

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Postgres struct {
	*xorm.Engine
}

func dbConfig() string {
	return fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		StringEnvVariable("DB_HOST"),
		StringEnvVariable("DB_PORT"),
		StringEnvVariable("DB_USER"),
		StringEnvVariable("DB_PASSWORD"),
		StringEnvVariable("DB_DATABASE"),
		StringEnvVariable("DB_SSLMODE"),
	)
}

func InitDatabase() (db *Postgres, err error) {
	dbconn, err := xorm.NewEngine("postgres", dbConfig())
	if err != nil {
		log.Println("engine creation failed", err)
	}
	return &Postgres{dbconn}, nil
}
