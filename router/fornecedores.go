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
	{path: "/fornecedores/nome/:filter",
		method:  http.MethodGet,
		handler: controllers.ListFornecedores,
	},
	{path: "/fornecedores/:id",
		method:  http.MethodGet,
		handler: controllers.ByIdFornecedor,
	},
	{path: "/fornecedores",
		method:  http.MethodPost,
		handler: controllers.CreateFornecedor,
	},
	{path: "/fornecedores/:id",
		method:  http.MethodPut,
		handler: controllers.EditFornecedor,
	},
	{path: "/fornecedores/:id",
		method:  http.MethodDelete,
		handler: controllers.DeleteForncedor,
	},
}
