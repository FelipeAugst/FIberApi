package main

import (
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
	if err := migrations.MigrateTables(); err != nil {
		log.Fatal(err)
	}
	log.Fatal(app.Listen(":5000"))
}
