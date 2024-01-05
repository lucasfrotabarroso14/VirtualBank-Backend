package routes

import (
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/controllers"
	"net/http"
)

var AccountRoutes = []Route{
	{
		URI:      "/accounts",
		Method:   http.MethodGet,
		Function: controllers.GetAccounts,
		NeedAuth: false,
	},
	{
		URI:      "/accounts",
		Method:   http.MethodPost,
		Function: controllers.CreateAccountHandler,
		NeedAuth: false,
	},
}
