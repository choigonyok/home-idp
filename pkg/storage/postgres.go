package storage

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/google/uuid"
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

func (c *PostgresClient) CreateAdminUser(uid float64) error {
	adminRoleId := uuid.NewString()
	applicantRoleId := uuid.NewString()
	adminPolicyId := uuid.NewString()
	applicantPolicyId := uuid.NewString()

	if _, err := c.DB().Exec(`INSERT INTO roles (id, name) VALUES ('` + adminRoleId + `', 'admin')`); err != nil {
		return err
	}

	if _, err := c.DB().Exec(`INSERT INTO roles (id, name) VALUES ('` + applicantRoleId + `', 'applicant')`); err != nil {
		return err
	}

	if _, err := c.DB().Exec(`INSERT INTO policies (id, name, policy) VALUES ('` + adminPolicyId + `', 'admin', '` + getAdminPolicy() + `')`); err != nil {
		return err
	}

	if _, err := c.DB().Exec(`INSERT INTO policies (id, name, policy) VALUES ('` + applicantPolicyId + `', 'applicant', '` + getApplicantPolicy() + `')`); err != nil {
		return err
	}

	if _, err := c.DB().Exec(`INSERT INTO rolepolicymapping (role_id, policy_id) VALUES ('` + adminRoleId + `', '` + adminPolicyId + `')`); err != nil {
		return err
	}

	if _, err := c.DB().Exec(`INSERT INTO rolepolicymapping (role_id, policy_id) VALUES ('` + applicantRoleId + `', '` + applicantPolicyId + `')`); err != nil {
		return err
	}

	fmt.Println("INT64 TO FLOAT64:", uid)

	if _, err := c.DB().Exec(`INSERT INTO users (id, name, role_id) VALUES ('` + strconv.FormatFloat(uid, 'e', -1, 64) + `', '` + env.Get("HOME_IDP_ADMIN_GIT_USERNAME") + `', '` + adminRoleId + `')`); err != nil {
		return err
	}

	return nil
}

func getAdminPolicy() string {
	return `{
		"policy": {
			"effect": "Allow",
			"target": "*",
			"action": "*"
		}
	}`
}

func getApplicantPolicy() string {
	return `{
	"policy": {
		"effect": "Deny",
		"target": "*",
		"action": "*"
	}
}`
}

func (c *PostgresClient) GetQueryFromProjects(cols ...string) (*sql.Row, error) {

	return nil, nil
}

// {
// 	"policy": {
// 		"effect": "Deny",
// 		"target": ["projects"],
// 		"action": ["CREATE"]
// 	}
// }
