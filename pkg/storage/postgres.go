package storage

import (
	"database/sql"
	"fmt"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	secretManagerDefaultTableName = "secret"
)

type PostgreSQLClient struct {
	table       string
	client      *sql.DB
	putQuery    string
	getQuery    string
	deleteQuery string
	listQuery   string
}

func (pc *PostgreSQLClient) Close() error {
	return pc.client.Close()
}

func (pc *PostgreSQLClient) Init(component config.Components) error {
	switch component {
	case 0: // Secret-Manager
		pc.initSecretManagerPostgreSQL()
	}
	return nil
}

func newPostgreSQLClient(component config.Components) *PostgreSQLClient {
	url := getPostgresqlUrl(component)
	db, err := sql.Open("pgx", url)
	fmt.Println("OPEN ERR:", err)
	fmt.Println("OPEN ERR:", err)
	fmt.Println("OPEN ERR:", err)

	return &PostgreSQLClient{
		client: db,
	}
}

func getPostgresqlUrl(component config.Components) string {
	prefix := ""
	switch component {
	case 0:
		prefix = "SECRET_MANAGER"
	}

	username := env.Get(prefix + "_STORAGE_USERNAME")
	password := env.Get(prefix + "_STORAGE_PASSWORD")
	host := env.Get(prefix + "_STORAGE_HOST")
	port := env.Get(prefix + "_STORAGE_PORT")
	database := env.Get(prefix + "_STORAGE_DATABASE")

	fmt.Printf("URL: postgres://%s:%s@%s:%s/%s", username, password, host, port,
		database)
	fmt.Printf("URL: postgres://%s:%s@%s:%s/%s", username, password, host, port,
		database)
	fmt.Printf("URL: postgres://%s:%s@%s:%s/%s", username, password, host, port,
		database)

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
}

func (pc *PostgreSQLClient) initSecretManagerPostgreSQL() error {
	if err := initializeSecretManagerPostgrSQLTables(pc.client); err != nil {
		return err
	}
	if err := initializeSecretManagerPostgreSQLFuncions(pc.client); err != nil {
		return err
	}
	pc.getQuery = getSecretManagerPostgreSQLQuery("get")
	pc.listQuery = getSecretManagerPostgreSQLQuery("list")
	pc.deleteQuery = getSecretManagerPostgreSQLQuery("delete")
	pc.putQuery = getSecretManagerPostgreSQLQuery("put")

	return nil
}

// createInitialTables create initial tables for components
func initializeSecretManagerPostgrSQLTables(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE ` + secretManagerDefaultTableName + ` (
	secret_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	project VARCHAR(100) NOT NULL,
	value_hash TEXT NOT NULL,
	key_hash VARCHAR(100) NOT NULL,
);
	`)
	return err
}

func initializeSecretManagerPostgreSQLFuncions(db *sql.DB) error {
	// 	_, err := db.Exec(`
	// CREATE OR REPLACE FUNCTION get_user_by_id(user_id INT)
	// RETURNS TABLE(id INT, username TEXT, email TEXT) AS $$
	// BEGIN
	// 		RETURN QUERY
	// 		SELECT id, username, email FROM users WHERE id = user_id;
	// END;
	// $$ LANGUAGE plpgsql;
	// 	`)
	// 	return err
	return nil
}

func getSecretManagerPostgreSQLQuery(method string) string {
	switch method {
	case "put":
		return "INSERT INTO " + secretManagerDefaultTableName + " VALUES($1, $2, $3, $4)" + " ON CONFLICT (path, key) DO " + " UPDATE SET (project, path, key, value) = ($1, $2, $3, $4)"
	case "get":
		return "SELECT value FROM " + secretManagerDefaultTableName + " WHERE path = $1 AND key = $2"
	case "delete":
		return "DELETE FROM " + secretManagerDefaultTableName + " WHERE path = $1 AND key = $2"
	case "list":
		return "SELECT key FROM " + secretManagerDefaultTableName + " WHERE path = $1" + " UNION ALL SELECT DISTINCT substring(substr(path, length($1)+1) from '^.*?/') FROM " + secretManagerDefaultTableName + " WHERE parent_path LIKE $1 || '%'"
	}
	return ""
}
