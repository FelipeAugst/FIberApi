package router

import (
	"net/http"
)

var routeFornecedores = []Route{
	{path: "/fornecedores",
		method:  http.MethodGet,
		handler: nil,
	},
	{path: "/fornecedores/:filter",
		method:  http.MethodGet,
		handler: nil,
	},
	{path: "/fornecedores",
		method:  http.MethodPost,
		handler: nil,
	},
	{path: "/fornecedores/:id",
		method:  http.MethodPut,
		handler: nil,
	},
	{path: "/fornecedores/:id",
		method:  http.MethodDelete,
		handler: nil,
	},
}
