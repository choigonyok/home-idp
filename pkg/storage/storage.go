package storage

import (
	"database/sql"
)

type StorageClient interface {
	Close() error
	DB() *sql.DB
	IsHealthy() bool
	Set(interface{})
	CreateAdminUser(username string) error
}
