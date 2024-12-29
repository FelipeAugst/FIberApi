package router

import (
	"api/controllers"
	"net/http"
)

func routeFornecedores(f controllers.Fornecedor) []Route {
	return []Route{
		{
			path:    "/fornecedores",
			method:  http.MethodGet,
			handler: f.ListAllFornecedores,
		},
		{
			path:    "/fornecedores/search/",
			method:  http.MethodGet,
			handler: f.SearchFornecedores,
		},
		{
			path:    "/fornecedores/:id",
			method:  http.MethodGet,
			handler: f.FindFornecedor,
		},
		{
			path:    "/fornecedores/create",
			method:  http.MethodPost,
			handler: f.CreateFornecedor,
		},
		{
			path:    "/fornecedores/:id/edit",
			method:  http.MethodPut,
			handler: f.EditFornecedor,
		},
		{
			path:    "/fornecedores/:id/delete",
			method:  http.MethodDelete,
			handler: f.DeleteForncedor,
		},
	}
}
