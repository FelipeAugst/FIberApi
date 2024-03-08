package router

import (
	"api/controllers"
	"net/http"
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
		path:    "/clientes/:id",
		method:  http.MethodPut,
		handler: controllers.EditCliente,
	},

	{
		path:    "/clientes/:id",
		method:  http.MethodDelete,
		handler: controllers.DeleteCliente,
	},
}
