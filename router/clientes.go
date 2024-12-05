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
		path:    "/clientes/:id",
		method:  http.MethodGet,
		handler: controllers.GetCliente,
	},
	{
		path:    "/clientes",
		method:  http.MethodGet,
		handler: controllers.GetAllCliente,
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
