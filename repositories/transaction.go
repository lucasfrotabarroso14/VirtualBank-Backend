package repositories

import (
	"database/sql"
	"fmt"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/models"
)

type Transaction struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *Transaction { return &Transaction{db} }

func (repository *Transaction) checkAccountExists(tx *sql.Tx, ID uint64) (bool, error) {

	statement, err := tx.Prepare("select count(*) from wallet where id_account=?")
	if err != nil {
		return false, err
	}

	result, err := statement.Exec(ID)
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return count >= 0, nil

}

func (repository Transaction) updateAccountBalance(tx *sql.Tx, accountID uint64, deltaAmount float64) error {
	_, err := tx.Exec("UPDATE wallet SET current_balance = current_balance + ? WHERE id_account = ?", deltaAmount, accountID)
	return err
}

func (repository *Transaction) checkSufficientFunds(tx *sql.Tx, accountID uint64, amount float64) error {
	var currentBalance float64
	err := tx.QueryRow("select current_balance FROM wallet WHERE id_account=?", accountID).Scan(&currentBalance)
	if err != nil {
		return err
	}

	if currentBalance < amount {
		return fmt.Errorf("insufficient funds in account %d", accountID)
	}
	return nil

}

func (repository Transaction) MakeTransaction(transaction models.Transaction) (uint64, error) {

	tx, err := repository.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// ver se existe a aconta de destino
	accountExists, err := repository.checkAccountExists(tx, transaction.IDdestinationAccount)
	if err != nil {
		return 0, err
	}
	if !accountExists {
		return 0, fmt.Errorf("Destination account with id %d does not exist", transaction.IDdestinationAccount)
	}

	if err := repository.checkSufficientFunds(tx, transaction.IDoriginAccount, transaction.Amount); err != nil { //aqui
		return 0, err
	}
	//vai remover o saldo da conta de origem
	if err := repository.updateAccountBalance(tx, transaction.IDoriginAccount, -transaction.Amount); err != nil {
		return 0, err
	}

	//vai adicionar o saldo  na conta destino
	if err := repository.updateAccountBalance(tx, transaction.IDdestinationAccount, transaction.Amount); err != nil {
		return 0, err
	}

	// registra a transacao
	statement, err := repository.db.Prepare(
		"INSERT  INTO  transactions (id_origin_account, transaction_type, amount, description, category, id_destination_account) VALUES (?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(
		transaction.IDoriginAccount, transaction.TransactionType,
		transaction.Amount, transaction.Description, transaction.Category, transaction.IDdestinationAccount,
	)
	if err != nil {
		return 0, err
	}

	// Commita a transação se tudo ocorrer sem erros
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	// Retorna o ID da transação inserida
	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil
}
