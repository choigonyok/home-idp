package storage

import "database/sql"

type PostgresClient struct {
	table       string
	client      *sql.DB
	putQuery    string
	getQuery    string
	deleteQuery string
	listQuery   string
}

func initPostgresTables(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE users (
	user_id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) NOT NULL,
	role_id INT NOT NULL
);

CREATE TABLE roles (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	policy JSON NOT NULL,
	project VARCHAR(100) NOT NULL
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
