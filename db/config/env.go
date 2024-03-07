package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DbName string
	DbPass string
)

func LoadEnvVArs() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env vars")
		SetEnvVars()

	}
}

func SetEnvVars() {
	DbName = os.Getenv("DbName")
	DbPass = os.Getenv("DbPass")

}
