package main

import (
	"api/config"
	"api/migrations"
	"api/router"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	go func() {
		signal.Notify(shutdown, os.Interrupt)
		signal.Notify(shutdown, syscall.SIGTERM)
	}()

	config.LoadEnvVArs()
	if err := migrations.MigrateTables(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{

		EnablePrintRoutes: true,
	})
	go func() {
		router.ConfigRoutes(app)
		app.Listen(config.PORT)
	}()
	<-shutdown

	log.Fatal("Shutdown")
	app.Shutdown()
}
