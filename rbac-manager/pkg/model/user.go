package model

import (
	"time"
)

type User struct {
	ID         int
	Name       string
	RoleID     int
	CreateTime time.Time
}
