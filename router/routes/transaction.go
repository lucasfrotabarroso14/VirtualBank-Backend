package routes

import (
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/controllers"
	"net/http"
)

var TransactionRoute = []Route{
	{
		URI:      "/transaction",
		Method:   http.MethodPost,
		Function: controllers.MakeTransactionHandler,
		NeedAuth: false,
	},
}
