package role

import (
	"fmt"
	"os"

	"github.com/choigonyok/home-idp/pkg/util"
)

type Action int

const (
	CREATE Action = iota
	READ
	DELETE
	UPDATE
	LIST
)

type Capability struct {
	ResourceName string
	Action       *Action
}

type Role struct {
	Name         string
	Token        string
	Capabilities []*Capability
}

type User struct {
	Name        string
	Roles       []*Role
	Email       string
	GitUsername string
}

type Project struct {
	Name  string
	Users []*User
	Admin *User
}

func setAdminUser() {
	username := os.Getenv("RBAC_ADMIN_USERNAME")
	pwKey := util.Hash(os.Getenv("RBAC_ADMIN_PASSWORD"))
	fmt.Println(username, pwKey)
}

func Get()      {}
func New()      {}
func Validate() {}
