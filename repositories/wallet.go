package repositories

import (
	"database/sql"
	"fmt"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/models"
)

type Wallet struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *Wallet { return &Wallet{db} }

func (repository Wallet) GetCurrentBalance(accountID uint64) (models.Wallet, error) {
	line, err := repository.db.Query(`select a.name, w.current_balance
										from accounts a
										inner join wallet w on a.id_account = w.id_account
										where a.id_account=?
										`, accountID)
	if err != nil {
		return models.Wallet{}, err
	}
	defer line.Close()
	var balance models.Wallet
	if line.Next() {
		if err = line.Scan(
			&balance.UserName,
			&balance.Current_balance,
		); err != nil {
			return models.Wallet{}, err
		}
	}
	return balance, nil

}

func (repository Wallet) CreateWallet(accountID uint64) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}

	}()

	result, err := tx.Exec(`INSERT INTO wallet (id_account,current_balance) values (?,0.0)`, accountID)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	// Verifica o número de linhas afetadas
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Verifica se pelo menos uma linha foi afetada
	if rowsAffected == 0 {
		return fmt.Errorf("falha ao criar a carteira para o usuário %d", accountID)
	}

	return nil

}
