package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/config"
)

func ConnectDB() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConnectionStringDB)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro

	}
	return db, nil
}
