package config

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
