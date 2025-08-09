package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var Database DBConfig
var SALT_ROUNDS int
var SECRET_KEY string

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Database = DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Name:     os.Getenv("MYSQL_DATABASE"),
	}
	SALT_ROUNDS, _ = strconv.Atoi(os.Getenv("SALT_ROUNDS"))
	SECRET_KEY = os.Getenv("SECRET_KEY")
}
