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
		"INSERT INTO goals (id_account, name,  icon_name, current_progress, goal_number, expected_date) values (?,?,?,?,?,?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	result, erro := statement.Exec(goal.ID_account, goal.Name, goal.Icon_name, goal.Current_progress, goal.Goal_number, goal.Expected_date)
	if erro != nil {
		return 0, erro
	}
	lastIdInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(lastIdInserted), nil
}

func (repository Goals) ListGoals(userID uint64) ([]models.Goal, error) {
	line, erro := repository.db.Query(`select id_goal, name, icon_name, current_progress, goal_number, expected_date from goals where id_account = ?`, userID)
	if erro != nil {
		return nil, erro
	}
	defer line.Close()
	var goals []models.Goal
	for line.Next() {
		var goal models.Goal
		if erro = line.Scan(
			&goal.ID_goal,
			&goal.Name,
			&goal.Icon_name,
			&goal.Current_progress,
			&goal.Goal_number,
			&goal.Expected_date,
		); erro != nil {
			return nil, erro
		}
		goals = append(goals, goal)
	}
	return goals, nil

}

func (repository Goals) GetByID(goal_id uint64) (models.Goal, error) {
	line, erro := repository.db.Query(`select * from goals where id_goal= ?`, goal_id)
	if erro != nil {
		return models.Goal{}, erro
	}
	defer line.Close()
	var goal models.Goal
	if line.Next() {
		if erro = line.Scan(
			&goal.ID_goal,
			&goal.ID_account,
			&goal.Name,
			&goal.Icon_name,
			&goal.Goal_number,
			&goal.Expected_date,
			&goal.Created_at,
			&goal.Updated_at,
		); erro != nil {
			return models.Goal{}, erro
		}
	}
	return goal, nil
}

func (repository Goals) DeleteGoal(goal_id uint64) error {
	statement, erro := repository.db.Prepare("delete from goals where id_goal = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(goal_id); erro != nil {
		return erro
	}
	return nil
}

func (repository Goals) UpdatedGoal(id_goal uint64, goal models.Goal) error {
	statement, erro := repository.db.Prepare("update goals set name = ?, icon_name = ?, current_progress = ?,goal_number = ? ")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(goal.Name, goal.Icon_name, goal.Current_progress, goal.Goal_number); erro != nil {
		return erro
	}
	return nil
}
