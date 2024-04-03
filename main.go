//go:build windows

package main

import (
	"api/config"
	"api/migrations"
	"api/router"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{

		EnablePrintRoutes: true,
	})
	config.LoadEnvVArs()
	router.ConfigRoutes(app)
	fmt.Println(config.DbName)
	if err := migrations.MigrateTables(); err != nil {
		log.Fatal(err)
	}
	log.Fatal(app.Listen(":5000"))
}
