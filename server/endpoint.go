package server

import (
	"biblioteca/controllers"
	"net/http"
)

var EndPoints = []EndPoint{
	{
		URI:      "/books",
		Method:   http.MethodPost,
		Function: controllers.RegisterBook,
		Auth:     false,
	},
	{
		URI:      "/books",
		Method:   http.MethodGet,
		Function: controllers.SearchAllBooks,
		Auth:     false,
	},
	{
		URI:      "/books/{id}",
		Method:   http.MethodGet,
		Function: controllers.SearchBookName,
		Auth:     false,
	},
	{
		URI:      "/books/{id}",
		Method:   http.MethodPut,
		Function: controllers.AlterBook,
		Auth:     false,
	},
	{
		URI:      "/books/{id}",
		Method:   http.MethodDelete,
		Function: controllers.ExcludeBook,
		Auth:     false,
	},
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.RegisterUser,
		Auth:     false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.SearchAllUsers,
		Auth:     false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodGet,
		Function: controllers.SearchUserID,
		Auth:     false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodPut,
		Function: controllers.AlterUser,
		Auth:     false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodDelete,
		Function: controllers.ExcludeUser,
		Auth:     false,
	},
}

