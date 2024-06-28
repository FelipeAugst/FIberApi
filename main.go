package main

import (
	"api/config"
	"api/migrations"
	"api/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnvVArs()
	if err := migrations.MigrateTables(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{

		EnablePrintRoutes: true,
	})

	router.ConfigRoutes(app)

	log.Fatal(app.Listen(config.PORT))
}
