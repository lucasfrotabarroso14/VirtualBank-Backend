package routes

import (
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/controllers"
	"net/http"
)

var loginRoute = Route{
	URI:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	NeedAuth: false,
}
