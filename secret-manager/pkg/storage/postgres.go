package storage

import "database/sql"

func initPostgresTables(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE secrets (
	user_id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) NOT NULL,
	role_id INT NOT NULL
);
`)
	return err
}

func initPostgresFunctions(db *sql.DB) error {
	return nil
}

func getPutQuery() string {
	return `
	PUTQUERY
	`
}

func getGetQuery() string {
	return `
	GETQUERY
	`
}

func getDeleteQuery() string {
	return `
	DELETEQUERY
	`
}

func getListQuery() string {
	return `
	LISTQUERY
	`
}
