package leaderboard

import (
	"contest/utils"
	"database/sql"
	"fmt"
	"time"
)

type Leaderboard struct {
	db *sql.DB
}

// Creates a Leaderboard instance after checking the env var ENVSPACE to get which variables related to redis options to load.
// Returns a pointer to said Leaderboard and an error to relay various issues.
func Init(db *sql.DB) (*Leaderboard, error) {
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	l := Leaderboard{
		db,
	}

	users, err := l.RetrieveUsers()
	if err != nil {
		return nil, fmt.Errorf("error retrieving users: %w", err)
	}

	problems, err := l.RetrieveProblems()
	if err != nil {
		return nil, fmt.Errorf("error retrieving problems: %w", err)
	}

	if err := l.initProblemStatuses(problems, users); err != nil {
		return nil, fmt.Errorf("error initiating problem statuses: %w", err)
	}

	return &l, nil
}

func (l *Leaderboard) RetrieveUsers() ([]utils.User, error) {
	rows, err := l.db.Query("SELECT * FROM users")
	if err != nil {
		return []utils.User{}, err
	}

	defer rows.Close()

	var users []utils.User

	for rows.Next() {
		var user utils.User

		if err := rows.Scan(&user.ID, &user.Pseudo); err != nil {
			return []utils.User{}, fmt.Errorf("error scanning for a user: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return []utils.User{}, fmt.Errorf("error encountered during query enumeration: %w", err)
	}

	return users, nil
}

func (l *Leaderboard) RetrieveProblems() ([]int64, error) {
	rows, err := l.db.Query("SELECT id FROM problems")
	if err != nil {
		return []int64{}, err
	}

	defer rows.Close()

	var problems []int64

	for rows.Next() {
		var problem int64

		if err := rows.Scan(&problem); err != nil {
			return []int64{}, fmt.Errorf("error scanning for a problem: %w", err)
		}

		problems = append(problems, problem)
	}

	if err := rows.Err(); err != nil {
		return []int64{}, fmt.Errorf("error encountered during query enumeration: %w", err)
	}

	return problems, nil

}

func (l *Leaderboard) initProblemStatuses(problems []int64, users []utils.User) error {
	tx, err := l.db.Begin()

	if err != nil {
		return fmt.Errorf("error starting a transaction: %w", err)
	}

	startedAt := time.Now().Unix()

	for _, user := range users {
		for _, problem := range problems {

			if _, err := tx.Exec(`INSERT INTO problems_statuses 
			(user_id, problem_id, status, submitted_at)
			VALUES
			(?, ?, ?, ?)`, user.ID, problem, utils.NotAttempted, startedAt); err != nil {
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					return fmt.Errorf("error initiating a problem status: %w.could not rollback: %w", err, rollbackErr)
				}
				return fmt.Errorf("error initiating a problem status: %w", err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error comtting transaction: %w", err)
	}

	return nil
}
