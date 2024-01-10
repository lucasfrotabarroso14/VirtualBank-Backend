package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/security"
	"strings"
	"time"
)

type Account struct {
	ID_account     uint64    `json:"id_account,omitempty"`
	Name           string    `json:"name,omitempty"`
	Email          string    `json:"email,omitempty"`
	Password       string    `json:"password,omitempty"`
	Status         string    `json:"status,omitempty"`
	Created_at     time.Time `json:"created_at,omitempty"`
	Updated_at     time.Time `json:"updated_at,omitempty"`
	Contact_number string    `json:"contact_number,omitempty"`
	Profile_image  []byte    `json:"profile_image,omitempty"`
}

func (account *Account) Prepare(step string) error {
	if erro := account.validate(step); erro != nil {
		return erro
	}
	if erro := account.format(step); erro != nil {
		return erro
	}
	return nil
}

func (account *Account) validate(step string) error {
	if account.Name == "" {
		return errors.New("O Nome é obrigatório e não pode estar em branco")
	}

	if account.Email == "" {
		return errors.New("O Email é obrigatório e não pode estar em branco")
		if erro := checkmail.ValidateFormat(account.Email); erro != nil {
			return errors.New("o email inserido é inválido")
		}
	}
	if step == "register" && account.Password == "" {
		return errors.New("A Senha é obrigatória e não pode estar em branco")
	}
	if step == "register" && account.Contact_number == "" {
		return errors.New("O Telefone é obrigatório e não pode estar em branco")
	}
	return nil

}

func (account *Account) format(step string) error {
	account.Name = strings.TrimSpace(account.Name)
	account.Email = strings.TrimSpace(account.Email)

	if step == "register" {
		passwordWithHash, erro := security.Hash(account.Password)
		if erro != nil {
			return erro
		}
		account.Password = string(passwordWithHash)
	}
	return nil
}
