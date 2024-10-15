package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type PostgresClient struct {
	Table       string
	Client      *sql.DB
	PutQuery    string
	GetQuery    string
	DeleteQuery string
	ListQuery   string
}

func (c *PostgresClient) Set(i interface{}) {
	c.Client = parseStorageClientFromInterface(i).DB()
}

func parseStorageClientFromInterface(i interface{}) *PostgresClient {
	client := i.(*PostgresClient)
	return client
}

func (c *PostgresClient) DB() *sql.DB {
	return c.Client
}

func (c *PostgresClient) Close() error {
	return c.Client.Close()
}

func (c *PostgresClient) IsHealthy() bool {
	err := c.Client.Ping()
	if err != nil {
		fmt.Println("TEST POSTGRESQL HEALTHY ERR: ", err)
		return false
	}

	return true
}

func NewPostgresClient(username, password, database string) StorageClient {
	host := "home-idp-postgres-postgresql.idp-system.svc.cluster.local"
	port := 5432
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	fmt.Println("TEST POSTGRESQL INFORMATION : ", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("TEST CREATE POSTGRESQL CLIENT ERROR: ", err)
	}

	for {
		if db.Ping() == nil {
			break
		}
		fmt.Println("WAITING FOR POSTGRESQL DB RUNNING")
		time.Sleep(time.Second)
	}

	return &PostgresClient{
		Client: db,
	}
}
