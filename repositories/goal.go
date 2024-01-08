package repositories

import (
	"database/sql"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/models"
)

type Goals struct {
	db *sql.DB
}

func NewGoalRepository(db *sql.DB) *Goals { return &Goals{db} }

func (repository Goals) CreateGoal(goal models.Goal) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"INSERT INTO goals (name, id_account,  icon_name, current_progress, goal_number, expected_date) values (?,?,?,?,?,?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	result, erro := statement.Exec(goal.Name, goal.ID_account, goal.Icon_name, goal.Current_progress, goal.Goal_number, goal.Expected_date)
	if erro != nil {
		return 0, erro
	}
	lastIdInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(lastIdInserted), nil
}

func (repository Goals) ListGoals() (models.Goal, error) {
	line, erro := repository.db.Query(`select id_goal, name, icon_name, current_progress, goal_number, expected_date from goals`)
	if erro != nil {
		return models.Goal{}, erro
	}
	defer line.Close()
	var goal models.Goal
	if line.Next() {
		if erro = line.Scan(
			&goal.ID_goal,
			&goal.Name,
			&goal.Icon_name,
			&goal.Current_progress,
			&goal.Goal_number,
			&goal.Expected_date,
		); erro != nil {
			return models.Goal{}, erro
		}
	}
	return goal, nil

}
