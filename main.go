package main

import (
	"api/db/config"
	"api/migrations"
	"api/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{

		EnablePrintRoutes: true,
	})
	router.ConfigRoutes(app)
	config.LoadEnvVArs()
	if err := migrations.MigrateTables(); err != nil {
		log.Fatal(err)
	}
	log.Fatal(app.Listen(":5000"))
}
