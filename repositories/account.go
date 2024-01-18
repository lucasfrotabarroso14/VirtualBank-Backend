package repositories

import (
	"database/sql"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/models"
)

type Account struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *Account { return &Account{db} }

func (repository Account) GetUserInfo(accountID uint64) (models.Account, error) {
	lines, erro := repository.db.Query("select profile_image, name, email, contact_number from accounts where id_account=?",
		accountID,
	)
	if erro != nil {
		return models.Account{}, erro
	}
	defer lines.Close()

	var account models.Account

	if lines.Next() {
		if erro = lines.Scan(
			&account.Profile_image,
			&account.Name,
			&account.Email,
			&account.Contact_number,
		); erro != nil {
			return models.Account{}, erro
		}
	}
	return account, nil

}

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
			&account.ID_account,
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

func (repository Account) getAccountByID(accountID uint64) (models.Account, error) {
	line, erro := repository.db.Query(`select * from accounts where id_account = ?`, accountID)
	if erro != nil {
		return models.Account{}, erro
	}
	defer line.Close()
	var account models.Account
	if line.Next() {
		if erro = line.Scan(
			&account.ID_account,
			&account.Name,
			&account.Email,
			&account.Password,
			&account.Status,
			&account.Created_at,
			&account.Updated_at,
			&account.Contact_number,
			&account.Profile_image,
		); erro != nil {
			return models.Account{}, erro
		}
	}
	return account, nil

}

func (repository Account) CreateAccount(account models.Account) (uint64, error) {
	statement, erro := repository.db.Prepare("insert into accounts (name,email,password, status, contact_number) values (?,?,?,?,?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()
	account.Status = "ATIVO"
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
	line, erro := repository.db.Query("select id_account,password from accounts where email = ?", email)

	if erro != nil {
		return models.Account{}, erro
	}

	defer line.Close()

	var account models.Account

	if line.Next() {
		if erro = line.Scan(&account.ID_account, &account.Password); erro != nil {
			return models.Account{}, erro
		}
	}
	return account, nil
}

// No arquivo repositories/account.go

func (repository Account) UpdateAccount(account models.Account) error {
	// Verifique se a conta existe
	existingAccount, err := repository.getAccountByID(account.ID_account)
	if err != nil {
		return err
	}

	// Atualize apenas os campos fornecidos
	if account.Name != "" {
		existingAccount.Name = account.Name
	}
	if account.Email != "" {
		existingAccount.Email = account.Email
	}
	if account.Status != "" {
		existingAccount.Status = account.Status
	}
	if account.Contact_number != "" {
		existingAccount.Contact_number = account.Contact_number
	}
	if len(account.Profile_image) > 0 {
		existingAccount.Profile_image = account.Profile_image
	}

	// Atualize a conta no banco de dados
	_, err = repository.db.Exec(`
		UPDATE accounts SET 
			name = ?,
			email = ?,
			status = ?,
			contact_number = ?,
			profile_image = ?
		WHERE id_account = ?
	`, existingAccount.Name, existingAccount.Email, existingAccount.Status, existingAccount.Contact_number, existingAccount.Profile_image, existingAccount.ID_account)

	return err
}
