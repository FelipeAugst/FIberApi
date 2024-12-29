package main

import (
	"api/config"
	"api/controllers"
	"api/migrations"
	"api/repository"
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

	fornecedorRepo, err := repository.NewFornecedorRepo()
	if err != nil {
		log.Fatal(err)
	}

	fornecedorController := controllers.NewFornecedor(&fornecedorRepo)

	go func() {
		router.ConfigRoutes(app, fornecedorController)
		app.Listen(config.PORT)
	}()
	<-shutdown

	log.Fatal("Shutdown")
	app.Shutdown()
}
