package router

import (
	"api/controllers"
	"net/http"
)

var routeVendas = []Route{
	{
		path:    "/vendas",
		method:  http.MethodGet,
		handler: controllers.ListAllVendas,
	},
	{
		path:    "/vendas/create",
		method:  http.MethodPost,
		handler: controllers.CreateVenda,
	},

	{
		path:    "/vendas/:id/update",
		method:  http.MethodPut,
		handler: controllers.UpdateVenda,
	},

	{
		path:    "/vendas/:id",
		method:  http.MethodGet,
		handler: controllers.FindVenda,
	},

	{
		path:    "/vendas/:id/delete",
		method:  http.MethodDelete,
		handler: controllers.DeleteVenda,
	},
	{
		path:    "/vendas/:id/conclude",
		method:  http.MethodDelete,
		handler: controllers.ConcludeVenda,
	},
}
