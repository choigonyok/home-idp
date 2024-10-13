package storage

import (
	"database/sql"
)

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
