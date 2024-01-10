package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/auth"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/database"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/models"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/repositories"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/responses"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateGoalHandler(w http.ResponseWriter, r *http.Request) {
	userID, erro := auth.ExtractUserID(r)

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
	//aqui abaixo ele vai mandar para a propriedade id do struct goal o id extraido no jwt
	goal.ID_account = userID

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
	userID, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewGoalRepository(db)
	goals, erro := repository.ListGoals(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, goals)

}

func GetGoalByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	goal_id, erro := strconv.ParseUint(params["id_goal"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewGoalRepository(db)
	goal, erro := repository.GetByID(goal_id)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, goal)

}

func DeleteGoalHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	goal_ID, erro := strconv.ParseUint(params["id_goal"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	db, erro := database.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	reporitory := repositories.NewGoalRepository(db)
	if erro = reporitory.DeleteGoal(goal_ID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)

	//goal, erro := reporitory.GetByID(goal_ID)
	//if erro != nil {
	//	responses.JSON(w, http.StatusInternalServerError, erro)
	//	return
	//}
	//if goal.ID_goal != goal_ID

}

func UpdateGoalHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	goal_id, erro := strconv.ParseUint(params["id_goal"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewGoalRepository(db)
	//goalSaved, erro := repository.GetByID(goal_id)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	//if goalSaved.ID_goal != goal_id {
	//	responses.Erro(w, http.StatusForbidden, errors.New("Não é possivel atualizar uma meta que nao seja sua"))
	//	return
	//}

	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}
	var goal models.Goal
	if erro = json.Unmarshal(requestBody, &goal); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}
	if goal.Prepare(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = repository.UpdatedGoal(goal_id, goal); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, goal)

}
