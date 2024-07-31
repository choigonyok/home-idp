package storage

import (
	"database/sql"
)

func initPostgresTables(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE project (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) NOT NULL,
	password_hash VARCHAR(100) NOT NULL,
	project_id INTEGER NOT NULL,
	FOREIGN KEY (project_id) REFERENCES project(id)
);

CREATE TABLE role (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	project_id INTEGER NOT NULL,
	FOREIGN KEY (project_id) REFERENCES project(id)
);

CREATE TABLE userrolemapping (
	user_id INTEGER NOT NULL,
	role_id INTEGER NOT NULL,
	PRIMARY KEY (user_id, role_id),
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (role_id) REFERENCES role(id)
);

CREATE TABLE policy (
	id SERIAL PRIMARY KEY,
	policy JSON NOT NULL,
	role_id INTEGER NOT NULL,
	FOREIGN KEY (role_id) REFERENCES role(id)
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
