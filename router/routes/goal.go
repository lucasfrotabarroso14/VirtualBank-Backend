package routes

import (
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/controllers"
	"net/http"
)

var GoalRoutes = []Route{
	{
		URI:      "/goals",
		Method:   http.MethodGet,
		Function: controllers.GetGoalsHandler,
		NeedAuth: false,
	},
	{
		URI:      "/goals",
		Method:   http.MethodPost,
		Function: controllers.CreateGoalHandler,
		NeedAuth: false,
	},
	{
		URI:      "/goals/{id_goal}",
		Method:   http.MethodGet,
		Function: controllers.GetGoalByIDHandler,
		NeedAuth: false,
	},
	{
		URI:      "/goals/{id_goal}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteGoalHandler,
		NeedAuth: false,
	},
	{
		URI:      "/goals/{id_goal}",
		Method:   http.MethodPut,
		Function: controllers.UpdateGoalHandler,
		NeedAuth: false,
	},
}
