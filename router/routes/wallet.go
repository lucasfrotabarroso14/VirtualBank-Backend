package routes

import (
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/controllers"
	"net/http"
)

var WalletRoutes = []Route{
	{
		URI:      "/wallet",
		Method:   http.MethodGet,
		Function: controllers.GetCurrentBalanceHandler,
		NeedAuth: false,
	},
}
