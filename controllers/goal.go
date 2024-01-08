package controllers

import (
	"encoding/json"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/database"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/models"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/repositories"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/responses"
	"io/ioutil"
	"net/http"
)

func CreateGoalHandler(w http.ResponseWriter, r *http.Request) {

	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return

	}
	var goal models.Goal
	if erro = json.Unmarshal(requestBody, &goal); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = goal.Prepare(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}
	db, erro := database.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewGoalRepository(db)
	goal.ID_goal, erro = repository.CreateGoal(goal)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusCreated, goal)

}

func GetGoalsHandler(w http.ResponseWriter, r *http.Request) {
	db, erro := database.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewGoalRepository(db)
	goals, erro := repository.ListGoals()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, goals)

}

//func DeleteGoalHandler(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	goalID, erro := strconv.ParseUint(params["id_goal"], 10, 64)
//	if erro != nil {
//		responses.Erro(w, http.StatusUnauthorized, erro)
//		return
//	}
//
//}
