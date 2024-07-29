package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
	_ "github.com/jackc/pgx/v4/stdlib"
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

func (pc *PostgreSQLClient) Init(component util.Components) error {
	switch component {
	case 0: // Secret-Manager
		log.Printf("Start initializing postgresql database...")
		pc.initSecretManagerPostgreSQL()
	case 2: // Rbac-Manager
		log.Printf("Start initializing postgresql database...")
		pc.initRbacManagerPostgreSQL()
	}
	return nil
}

func newPostgreSQLClient(component util.Components) *PostgreSQLClient {
	url := getPostgresqlUrl(component)
	log.Printf("Start connecting with postgresql database...")
	db, _ := sql.Open("pgx", url)

	return &PostgreSQLClient{
		client: db,
	}
}

func getPostgresqlUrl(component util.Components) string {
	prefix := env.GetEnvPrefix(component)

	username := env.Get(prefix + "_MANAGER_STORAGE_USERNAME")
	password := env.Get(prefix + "_MANAGER_STORAGE_PASSWORD")
	host := env.Get(prefix + "_MANAGER_STORAGE_HOST")
	port := env.Get(prefix + "_MANAGER_STORAGE_PORT")
	database := env.Get(prefix + "_MANAGER_STORAGE_DATABASE")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
}

func (pc *PostgreSQLClient) initSecretManagerPostgreSQL() error {
	log.Printf("--- Start creating initial postgresql database tables")
	if err := initializeSecretManagerPostgrSQLTables(pc.client); err != nil {
		return fmt.Errorf("\nFailed to create initial postgresql database tables")
	}

	log.Printf("--- Start creating initial postgresql database functions")
	if err := initializeSecretManagerPostgreSQLFuncions(pc.client); err != nil {
		return fmt.Errorf("\nFailed to create initial postgresql database functions")
	}
	pc.getQuery = getSecretManagerPostgreSQLQuery("get")
	pc.listQuery = getSecretManagerPostgreSQLQuery("list")
	pc.deleteQuery = getSecretManagerPostgreSQLQuery("delete")
	pc.putQuery = getSecretManagerPostgreSQLQuery("put")

	return nil
}

func (pc *PostgreSQLClient) initRbacManagerPostgreSQL() error {
	log.Printf("--- Start creating initial postgresql database tables")
	if err := initializeRbacManagerPostgrSQLTables(pc.client); err != nil {
		return fmt.Errorf("\nFailed to create initial postgresql database tables")
	}

	log.Printf("--- Start creating initial postgresql database functions")
	if err := initializeRbacManagerPostgreSQLFuncions(pc.client); err != nil {
		return fmt.Errorf("\nFailed to create initial postgresql database functions")
	}
	pc.getQuery = getRbacManagerPostgreSQLQuery("get")
	pc.listQuery = getRbacManagerPostgreSQLQuery("list")
	pc.deleteQuery = getRbacManagerPostgreSQLQuery("delete")
	pc.putQuery = getRbacManagerPostgreSQLQuery("put")

	return nil
}

// createInitialTables create initial tables for components
func initializeSecretManagerPostgrSQLTables(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE secrets (
	secret_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	project VARCHAR(100) NOT NULL,
	value_hash TEXT NOT NULL,
	key_hash VARCHAR(100) NOT NULL
);`)
	return err
}

// createInitialTables create initial tables for components
func initializeRbacManagerPostgrSQLTables(db *sql.DB) error {
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

func initializeRbacManagerPostgreSQLFuncions(db *sql.DB) error {
	return nil
}

func getSecretManagerPostgreSQLQuery(method string) string {
	switch method {
	case "put":
		return "INSERT INTO secrets VALUES($1, $2, $3, $4)" + " ON CONFLICT (path, key) DO " + " UPDATE SET (project, path, key, value) = ($1, $2, $3, $4)"
	case "get":
		return "SELECT value FROM secrets WHERE path = $1 AND key = $2"
	case "delete":
		return "DELETE FROM secrets WHERE path = $1 AND key = $2"
	case "list":
		return "SELECT key FROM secrets WHERE path = $1" + " UNION ALL SELECT DISTINCT substring(substr(path, length($1)+1) from '^.*?/') FROM secrets WHERE parent_path LIKE $1 || '%'"
	}
	return ""
}

func getRbacManagerPostgreSQLQuery(method string) string {
	switch method {
	case "put":
		return "INSERT INTO {TABLE_NAME} VALUES($1, $2, $3, $4)"
	case "get":
		return "SELECT value FROM {TABLE_NAME}"
	case "delete":
		return "DELETE FROM {TABLE_NAME}"
	case "list":
		return "SELECT key FROM {TABLE_NAME}"
	}
	return ""
}

// CREATE TABLE testtable (
// 	secret_id SERIAL PRIMARY KEY,
// 	user_id INT NOT NULL,
// 	project VARCHAR(100) NOT NULL,
// 	value_hash TEXT NOT NULL,
// 	key_hash VARCHAR(100) NOT NULL
// );
