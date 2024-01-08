package models

import (
	"errors"
	"strings"
	"time"
)

type Goal struct {
	ID_goal          uint64    `json:"id_goal,omitempty"`
	ID_account       uint64    `json:"id_account,omitempty"`
	Name             string    `json:"name,omitempty"`
	Icon_name        string    `json:"icon_name,omitempty"`
	Current_progress float64   `json:"current_progress,omitempty"`
	Goal_number      float64   `json:"goal_number,omitempty"`
	Expected_date    time.Time `json:"expected_date,omitempty"`
	Created_at       time.Time `json:"created_at,omitempty"`
	Updated_at       time.Time `json:"updated_at,omitempty"`
}

func (goal *Goal) Prepare() error {
	if erro := goal.validate(); erro != nil {
		return erro
	}
	goal.formate()

	return nil

}

func (goal *Goal) validate() error {
	if goal.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	return nil

}

func (goal *Goal) formate() {
	goal.Name = strings.TrimSpace(goal.Name)
}
