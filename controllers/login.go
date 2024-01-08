package controllers

import (
	"encoding/json"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/auth"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/database"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/models"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/repositories"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/responses"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/security"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var account models.Account

	if erro = json.Unmarshal(requestBody, &account); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewAccountRepository(db)
	AccountSavedOnDatabase, erro := repository.FindAccountByEmail(account.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if erro = security.VerifiyPassword(AccountSavedOnDatabase.Password, account.Password); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	token, erro := auth.CreateToken(AccountSavedOnDatabase.ID_account)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JWTtoJSON(w, http.StatusOK, token)
}
