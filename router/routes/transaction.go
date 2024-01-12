package routes

import (
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/controllers"
	"net/http"
)

var TransactionRoute = []Route{
	{
		URI:      "/transactions",
		Method:   http.MethodPost,
		Function: controllers.MakeTransactionHandler,
		NeedAuth: false,
	},
	{
		URI:      "/transactions",
		Method:   http.MethodGet,
		Function: controllers.GetUserTransactionsHandler,
		NeedAuth: false,
	},
}
