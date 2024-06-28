package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/usuarios",
		Method:       http.MethodPost,
		Handler:      controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios",
		Method:       http.MethodGet,
		Handler:      controllers.GetUsers,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios/{id}",
		Method:       http.MethodGet,
		Handler:      controllers.GetUser,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios/{id}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios/{id}",
		Method:       http.MethodDelete,
		Handler:      controllers.DeleteUser,
		AuthRequired: false,
	},
}
