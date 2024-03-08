package router

import (
	"api/controllers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var routeClientes = []Route{
	{
		path:    "/clientes",
		method:  http.MethodPost,
		handler: controllers.CreateCliente,
	},
	{
		path:    "/clientes/:filter",
		method:  http.MethodGet,
		handler: controllers.ListClientes,
	},
	{
		path:    "/clientes",
		method:  http.MethodGet,
		handler: controllers.ListAllClientes,
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
