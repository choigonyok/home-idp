package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	defaultTableName = "rbac"
)

type PostgresClient struct {
	Table       string
	Client      *sql.DB
	PutQuery    string
	GetQuery    string
	DeleteQuery string
	ListQuery   string
}

func (c *PostgresClient) DB() *sql.DB {
	return c.Client
}

func (c *PostgresClient) Close() error {
	return c.Client.Close()
}

func newPostgresDB(component util.Components) *sql.DB {
	url := getPostgresqlUrl(component)
	log.Printf("Start connecting with postgresql database...")
	db, _ := sql.Open("pgx", url)
	return db
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

// func initializeSecretManagerPostgrSQLTables(db *sql.DB) error {
// 	_, err := db.Exec(`CREATE TABLE secrets (
// 	secret_id SERIAL PRIMARY KEY,
// 	user_id INT NOT NULL,
// 	project VARCHAR(100) NOT NULL,
// 	value_hash TEXT NOT NULL,
// 	key_hash VARCHAR(100) NOT NULL
// );`)
// 	return err
// }
