package controllers

import (
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/auth"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/database"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/repositories"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/responses"
	"net/http"
)

func GetCurrentBalanceHandler(w http.ResponseWriter, r *http.Request) {
	accountID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	db, err := database.ConnectDB()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewWalletRepository(db)
	balance, err := repository.GetCurrentBalance(accountID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, balance)

}
