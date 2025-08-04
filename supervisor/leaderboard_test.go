package leaderboard

import (
	"contest/utils"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/tursodatabase/go-libsql"
)

// For some deferring, all tests will take place within this testing scope.
func TestSuite(t *testing.T) {
	dir, err := os.MkdirTemp("", "temp_db")
	if err != nil {
		t.Fatalf("Error generating temp dir: %v", err)
	}

	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, "temp.db")
	if err := os.WriteFile(dbPath, []byte{}, 0666); err != nil {
		t.Fatalf("Error writing local db: %v", err)
	}

	tursoDBName := fmt.Sprintf("file:%s", dbPath)
	db, err := sql.Open("libsql", tursoDBName)
	if err != nil {
		t.Fatalf("Error opening the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatalf("Error pinging the database: %v", err.Error())
	}

	// Create tables
	if err := utils.CreateTables(db); err != nil {
		t.Fatalf("Error creating tables: %v", err)
	}

	testInit_OK(t, db)

}

func testInit_OK(t *testing.T, db *sql.DB) {
	if _, err := Init(db); err != nil {
		t.Fatalf("Error executing Init: %v", err)
	}
}
