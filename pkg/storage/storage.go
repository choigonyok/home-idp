package storage

import (
	"database/sql"
)

type StorageClient interface {
	Close() error
	DB() *sql.DB
	IsHealthy() bool
	Set(interface{})
	CreateAdminUser(float64) error

	// GetQueryFromProjects(cols ...string) (*sql.Rows, error)
	// GetQueryFromUsers(cols ...string) (*sql.Rows, error)
}
