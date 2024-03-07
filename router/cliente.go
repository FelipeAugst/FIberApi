package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var ClienteRoutes = []Route{
	{
		path:   "/clientes",
		method: http.MethodGet,
		handler: func(*fiber.Ctx) error {
			return nil
		},
	},
	{
		path:   "/clientes/:filter",
		method: http.MethodGet,
		handler: func(*fiber.Ctx) error {
			return nil
		},
	},
	{
		path:   "/clientes",
		method: http.MethodPost,
		handler: func(*fiber.Ctx) error {
			return nil
		},
	},
	{
		path:   "/clientes/:id",
		method: http.MethodPut,
		handler: func(*fiber.Ctx) error {
			return nil
		},
	},

	{
		path:   "/clientes/:id",
		method: http.MethodDelete,
		handler: func(*fiber.Ctx) error {
			return nil
		},
	},
}
