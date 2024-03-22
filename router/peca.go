package router

import (
	"api/controllers"
	"net/http"
)

var routePecas = []Route{
	{path: "/pecas",
		method:  http.MethodGet,
		handler: controllers.ListAllPecas,
	},
	{path: "/pecas/descricao/:filter",
		method:  http.MethodGet,
		handler: controllers.ListPecas,
	},
	{path: "/pecas",
		method:  http.MethodPost,
		handler: controllers.CreatePeca,
	},
	{path: "/pecas/:id",
		method:  http.MethodPut,
		handler: controllers.EditPeca,
	},
	{path: "/pecas/:id",
		method:  http.MethodGet,
		handler: controllers.ByIdPeca,
	},
	{path: "/pecas/:id",
		method:  http.MethodDelete,
		handler: controllers.DeletePeca,
	},
}
