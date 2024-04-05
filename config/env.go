package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DbName string
	DbPass string
	PORT   string
)

func SetEnvVars() {
	DbName = os.Getenv("DB_NAME")
	DbPass = os.Getenv("DbPass")
	PORT = os.Getenv("PORT")

}

func LoadEnvVArs() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load env vars,error: %s", err.Error())

	}
	SetEnvVars()
}
