package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/peopleig/food-ordering-go/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() (*sql.DB, error) {
	cfg := config.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	var err error
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %v", err)
	}
	fmt.Println("Database connected successfully!")
	return DB, nil
}

func CloseDatabase() error {
	if DB != nil {
		fmt.Println("Closing database connection...")
		return DB.Close()
	}
	return nil
}
