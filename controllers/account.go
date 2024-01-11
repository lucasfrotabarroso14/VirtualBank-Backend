package controllers

import (
	"encoding/json"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/auth"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/database"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/models"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/repositories"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/responses"
	"io/ioutil"
	"net/http"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	db, erro := database.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewAccountRepository(db)

	accounts, erro := repository.ListAccounts()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, accounts)

}
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	userID, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	db, erro := database.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	repository := repositories.NewAccountRepository(db)
	userInfo, erro := repository.GetUserInfo(userID)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	responses.JSON(w, http.StatusOK, userInfo)

}

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
	}
	var account models.Account

	if erro = json.Unmarshal(bodyRequest, &account); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = account.Prepare("register"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	accountRepository := repositories.NewAccountRepository(db)
	walletRepository := repositories.NewWalletRepository(db)

	account.ID_account, erro = accountRepository.CreateAccount(account)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}

	if erro := walletRepository.CreateWallet(account.ID_account); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, account)

}
