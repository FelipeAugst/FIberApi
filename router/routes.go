package router

import (
	"github.com/gofiber/fiber/v2"
)

func ConfigRoutes(app *fiber.App) {
	var routes []Route
	routes = append(routes, routePecas...)
	routes = append(routes, routeClientes...)
	routes = append(routes, routeFornecedores...)

	for _, route := range routes {
		app.Add(route.method, route.path, route.handler)
	}

}
