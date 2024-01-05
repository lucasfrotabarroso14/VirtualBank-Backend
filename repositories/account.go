package repositories

import (
	"database/sql"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/models"
)

type Account struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *Account { return &Account{db} }

func (repository Account) ListAccounts() ([]models.Account, error) {
	linhas, erro := repository.db.Query(`select * from accounts`)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	var accounts []models.Account
	for linhas.Next() {
		var account models.Account
		if erro = linhas.Scan(
			&account.Account_id,
			&account.Name,
			&account.Email,
			&account.Password,
			&account.Status,
			&account.Created_at,
			&account.Updated_at,
			&account.Contact_number,
			&account.Profile_image,
		); erro != nil {
			return nil, erro
		}

		accounts = append(accounts, account)
	}
	return accounts, nil

}

func (repository Account) CreateAccount(account models.Account) (uint64, error) {
	statement, erro := repository.db.Prepare("insert into accounts (name,email,password, status, contact_number) values (?,?,?,?,?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()
	result, erro := statement.Exec(account.Name, account.Email, account.Password, account.Status, account.Contact_number)

	if erro != nil {
		return 0, erro
	}

	lastIdInserted, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastIdInserted), nil

}

func (repository Account) FindAccountByEmail(email string) (models.Account, error) {
	line, erro := repository.db.Query("select account_id,password from accounts where email = ?", email)

	if erro != nil {
		return models.Account{}, erro
	}

	defer line.Close()

	var account models.Account

	if line.Next() {
		if erro = line.Scan(&account.Account_id, &account.Password); erro != nil {
			return models.Account{}, erro
		}
	}
	return account, nil
}
