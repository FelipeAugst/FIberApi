package router

import (
	"api/controllers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var routeVendas = []Route{
	{
		path:    "/vendas",
		method:  http.MethodGet,
		handler: controllers.ListVendas,
	},
	{
		path:    "/vendas/:cliente/:peca",
		method:  http.MethodPost,
		handler: controllers.CreateVenda,
	},

	{
		path:   "/vendas/:id",
		method: http.MethodPost,
		handler: func(*fiber.Ctx) error {
			return nil
		},
	},

	{
		path:   "/vendas/peca/:filter",
		method: http.MethodGet,
		handler: func(*fiber.Ctx) error {
			return nil
		},
	},
	{
		path:   "/vendas/fornecedor/:filter",
		method: http.MethodGet,
		handler: func(*fiber.Ctx) error {
			return nil
		},
	},
	{
		path:   "/vendas",
		method: http.MethodPut,
		handler: func(*fiber.Ctx) error {
			return nil
		},
	},

	{
		path:   "/vendas",
		method: http.MethodDelete,
		handler: func(*fiber.Ctx) error {
			return nil
		},
	},
}
