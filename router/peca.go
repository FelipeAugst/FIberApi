package router

import (
	"api/controllers"
	"net/http"
)

var routePecas = []Route{
	{path: "/pecas",
		method:  http.MethodGet,
		handler: controllers.ListaPecas,
	},
	{path: "/pecas/{desc}",
		method:  http.MethodGet,
		handler: controllers.BuscaPecas,
	},
	{path: "/pecas",
		method:  http.MethodPost,
		handler: controllers.CriaPeca,
	},
	{path: "/pecas",
		method:  http.MethodPut,
		handler: controllers.EditaPeca,
	},
	{path: "/pecas",
		method:  http.MethodDelete,
		handler: controllers.DeletaPeca,
	},
}
