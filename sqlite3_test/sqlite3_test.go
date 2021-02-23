package sqlite3_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func getDatabaseHandle(dbpath string) (*sql.DB, error) {
	database, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("Failed to create the handle")
		return nil, err
	}
	if err = database.Ping(); err != nil {
		fmt.Printf("Failed to keep connection alive")
		return nil, err
	}
	return database, nil
}

func getAllRows(database *sql.DB, table string) {
	query := fmt.Sprintf("SELECT User, AppName FROM %s LIMIT 10", table)
	rows, err := database.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var app string
		rows.Scan(&id, &app)
		fmt.Println(id, app)
	}
}

func TestSqlite3(t *testing.T) {
	db, err := getDatabaseHandle("./data/gozo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	getAllRows(db, "timesheet")
}