package utils

import (
	"database/sql"

	"github.com/go-faker/faker/v4"
)

const (
	USERS_NUMBER    int = 100
	PROBLEMS_NUMBER int = 5
)

func CreateTables(db *sql.DB) error {
	tablesHandlers := []func(*sql.DB) error{
		usersHandler,
		problemsHandler,
		problemsStatusesHandler,
	}

	for _, handler := range tablesHandlers {
		if err := handler(db); err != nil {
			return err
		}
	}

	return nil
}

func usersHandler(db *sql.DB) error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		pseudo TEXT
	)`); err != nil {
		return err
	}

	for range USERS_NUMBER {
		db.Exec(`INSERT INTO users (pseudo) VALUES (?)`, faker.Username())
	}

	return nil
}

func problemsHandler(db *sql.DB) error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS problems (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT
	)`); err != nil {
		return err
	}

	for range PROBLEMS_NUMBER {
		db.Exec(`INSERT INTO problems (name) VALUES (?)`, faker.Username())
	}

	return nil
}

func problemsStatusesHandler(db *sql.DB) error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS problems_statuses (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER REFERENCES users(id),
	problem_id INTEGER REFERENCES problems(id),
	status INTEGER,
	submitted_at INTEGER
	)`); err != nil {
		return err
	}

	return nil
}
