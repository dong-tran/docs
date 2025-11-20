package infrastructure

import (
"github.com/jmoiron/sqlx"
_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "./orders.db")
	if err != nil {
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS orders (
id TEXT PRIMARY KEY,
customer_id TEXT NOT NULL,
items TEXT NOT NULL,
total_amount REAL NOT NULL,
currency TEXT NOT NULL,
status TEXT NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL
);
	`

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
