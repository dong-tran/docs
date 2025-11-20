package infrastructure

import (
"github.com/jmoiron/sqlx"
_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "./tasks.db")
	if err != nil {
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
id INTEGER PRIMARY KEY AUTOINCREMENT,
title TEXT NOT NULL,
description TEXT,
completed BOOLEAN NOT NULL DEFAULT 0,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL
);
	`

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
