package data_access

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	// createUsersTable := `
	// CREATE TABLE IF NOT EXISTS users (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	username TEXT NOT NULL UNIQUE,
	// 	password TEXT NOT NULL,
	// 	email TEXT NOT NULL UNIQUE,
	// 	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	// );
	// `

	// _, err := DB.Exec(createUsersTable)
	// if err != nil {
	// 	panic(err)
	// }

	// createEventsTable := `
	// CREATE TABLE IF NOT EXISTS events (

	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	name TEXT NOT NULL,
	// 	description TEXT NOT NULL,
	// 	location TEXT NOT NULL,
	// 	date_time DATETIME NOT NULL,
	// 	user_id INTEGER NOT NULL,
	// 	FOREIGN KEY (user_id) REFERENCES users (id)
	// );
	// `

	// _, err = DB.Exec(createEventsTable)
	// if err != nil {
	// 	panic(err)
	// }

	// createRegistrationTable := `
	// CREATE TABLE IF NOT EXISTS registration (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	event_id INTEGER NOT NULL,
	// 	user_id INTEGER NOT NULL,
	// 	FOREIGN KEY (event_id) REFERENCES events (id),
	// 	FOREIGN KEY (user_id) REFERENCES users (id)
	// );
	// `

	// _, err := DB.Exec(createRegistrationTable)
	// if err != nil {
	// 	panic(err)
	// }

}
