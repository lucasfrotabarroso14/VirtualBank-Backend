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

func MakeTransactionHandler(w http.ResponseWriter, r *http.Request) {
	accountID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	RequestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var transaction models.Transaction
	if err = json.Unmarshal(RequestBody, &transaction); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	transaction.IDoriginAccount = accountID
	db, err := database.ConnectDB()

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewTransactionRepository(db)

	transaction.IDtransaction, err = repository.MakeTransaction(transaction)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.MakeJSONResponse(w, http.StatusCreated, "Transferência realizada com sucesso")

}
