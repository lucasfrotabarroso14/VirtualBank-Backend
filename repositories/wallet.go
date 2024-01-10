package repositories

import (
	"database/sql"
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
