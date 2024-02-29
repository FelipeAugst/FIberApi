package router

import "github.com/gofiber/fiber/v2"

type Route struct {
	path    string
	method  string
	handler func(*fiber.Ctx) error
}
