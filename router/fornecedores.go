package router

import (
	"api/controllers"
	"net/http"
)

var routeFornecedores = []Route{
	{path: "/fornecedores",
		method:  http.MethodGet,
		handler: controllers.ListAllFornecedores,
	},
	{path: "/fornecedores/search/",
		method:  http.MethodGet,
		handler: controllers.SearchFornecedores,
	},
	{path: "/fornecedores/:id",
		method:  http.MethodGet,
		handler: controllers.FindFornecedor,
	},
	{path: "/fornecedores/create",
		method:  http.MethodPost,
		handler: controllers.CreateFornecedor,
	},
	{path: "/fornecedores/:id/edit",
		method:  http.MethodPut,
		handler: controllers.EditFornecedor,
	},
	{path: "/fornecedores/:id/delete",
		method:  http.MethodDelete,
		handler: controllers.DeleteForncedor,
	},
}
