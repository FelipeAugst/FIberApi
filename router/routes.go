package router

import (
	"api/controllers"

	"github.com/gofiber/fiber/v2"
)

func ConfigRoutes(app *fiber.App, fornecedorController controllers.Fornecedor) {
	var routes []Route
	routes = append(routes, routePecas...)
	routes = append(routes, routeClientes...)
	routes = append(routes, routeFornecedores(fornecedorController)...)
	routes = append(routes, routeVendas...)
	routes = append(routes, vendaItem...)

	for _, route := range routes {
		app.Add(route.method, route.path, route.handler)
	}
}
