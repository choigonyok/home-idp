package system

import (
	"fmt"

	"github.com/choigonyok/idp/pkg/model"
	"github.com/choigonyok/idp/pkg/store"
)

type System struct {
	db *store.SQLiteStore
}

func NewSystem(db *store.SQLiteStore) *System {
	return &System{
		db: db,
	}
}

func (s *System) Create(m *model.System) error {
	switch m.Type {
	case "env":
		return s.db.CreateEnv(m)
	case "team":
		return s.db.CreateTeam(m)
	case "permission":
		return s.db.CreatePermission(m)
	case "role":
		return s.db.CreateRole(m)
	default:
		return fmt.Errorf("Invalid system type")
	}
}

func (s *System) List(t string) (*[]model.System, error) {
	switch t {
	case "envs":
		return s.db.ListEnvs()
	case "teams":
		return s.db.ListTeams()
	case "permissions":
		return s.db.ListPermissions()
	case "roles":
		return s.db.ListRoles()
	default:
		return nil, fmt.Errorf("Invalid system type")
	}
}

func (s *System) Update(m *model.System) error {
	switch m.Type {
	case "roles":
		return s.db.UpdateRole(m)
	default:
		return nil
	}
}

func (s *System) ListRolePermissions(roleID string) (*[]model.System, error) {
	return s.db.ListPermissionsByRoleID(roleID)
}
