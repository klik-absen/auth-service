package db

import (
	"fmt"
	"ka-auth-service/internal/infrastructure/env"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDBConnection(config *env.Config) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	return db
}
