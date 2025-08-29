package store

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/choigonyok/idp/pkg/model"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type SQLiteStore struct {
	db *sql.DB
}

type User struct {
	ID       string
	Username string
	Password string // hashed
}

func NewSQLiteStore(path string) *SQLiteStore {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("failed open db: %v", err)
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, username TEXT UNIQUE, password TEXT)`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS infra (id TEXT PRIMARY KEY, name TEXT UNIQUE, status TEXT, version TEXT NOT NULL CHECK (version <> ''), platform TEXT NOT NULL CHECK (platform <> ''), env TEXT NOT NULL CHECK (env <> ''), created_at  DATETIME DEFAULT (CURRENT_TIMESTAMP), assignee TEXT NOT NULL CHECK (assignee <> ''), context TEXT)`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS services (id TEXT PRIMARY KEY, name TEXT UNIQUE, status TEXT, version TEXT NOT NULL CHECK (version <> ''), image TEXT, tag TEXT, platform TEXT, namespace TEXT, env TEXT NOT NULL CHECK (env <> ''), created_at  DATETIME DEFAULT (CURRENT_TIMESTAMP), assignee TEXT NOT NULL CHECK (assignee <> ''), git TEXT, FOREIGN KEY (platform) REFERENCES infra(id) ON DELETE CASCADE )`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS envs (id TEXT PRIMARY KEY, name TEXT UNIQUE )`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS teams (id TEXT PRIMARY KEY, name TEXT UNIQUE, leader_id TEXT DEFAULT "", FOREIGN KEY (leader_id) REFERENCES users(id) )`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS permissions (id TEXT PRIMARY KEY, name TEXT UNIQUE, json TEXT NOT NULL, description TEXT)`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	// _, err = db.Exec(`DROP TABLE roles`)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS roles
	 (id TEXT PRIMARY KEY, name TEXT UNIQUE, description TEXT)`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS role_permissions ( role_id TEXT NOT NULL, permission_id TEXT NOT NULL, PRIMARY KEY(role_id, permission_id), FOREIGN KEY(role_id) REFERENCES roles(id), FOREIGN KEY(permission_id) REFERENCES permissions(id))`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	// _, err = db.Exec(`UPDATE infra SET context = 't1' WHERE name = 't1'`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err = db.Exec(`UPDATE infra SET context = 'kind3' WHERE name = 'kind3'`)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS metrics (id TEXT PRIMARY KEY, service_name TEXT UNIQUE, metric_name TEXT, source TEXT, query TEXT, active TEXT)`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	return &SQLiteStore{db: db}
}

func (s *SQLiteStore) Authenticate(username, password string) (*User, error) {
	row := s.db.QueryRow(`SELECT id, username, password FROM users WHERE username = ?`, username)
	var u User
	if err := row.Scan(&u.ID, &u.Username, &u.Password); err != nil {
		return nil, errors.New("no such user")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}
	return &u, nil
}

func timeNowUnix() int64 {
	return int64((int64)((int)((int64)(0))))
}

func (s *SQLiteStore) GetServiceNamespace(serviceID string) (string, error) {
	r := s.db.QueryRow(`SELECT namespace FROM services WHERE id = ?`, serviceID)
	ns := ""
	if err := r.Scan(&ns); err != nil {
		return "", err
	}
	return ns, nil
}

func (s *SQLiteStore) GetServiceNameFromServiceID(serviceID string) (string, error) {
	r := s.db.QueryRow(`SELECT name FROM services WHERE id = ?`, serviceID)
	name := ""
	if err := r.Scan(&name); err != nil {
		return "", err
	}
	return name, nil
}

func (s *SQLiteStore) GetServicePlatform(serviceID string) (string, error) {
	r := s.db.QueryRow(`SELECT platform FROM services WHERE id = ?`, serviceID)
	platform := ""
	if err := r.Scan(&platform); err != nil {
		return "", err
	}
	return platform, nil
}

func (s *SQLiteStore) CreateService(svc *model.Service) error {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec(`INSERT INTO services (id, name, status, version, image, tag, platform, namespace, env, assignee, git) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, hex.EncodeToString(b), svc.Name, svc.Status, svc.Version, svc.Image, svc.Tag, svc.PlatformID, svc.Namespace, svc.Env, svc.Assignee, svc.Git); err != nil {
		return err
	}

	for _, m := range svc.Metrics {
		if _, err := s.db.Exec(`INSERT INTO services (id, service_name, metric_name, source, query, active) VALUES (?, ?, ?, ?, ?)`, m.ID, m.ServiceName, m.MetricName, m.Source, m.Query, m.Active); err != nil {
			return err
		}
	}

	return nil
}

func (s *SQLiteStore) GetInfraFromContext(context string) (*model.Infra, error) {
	r := s.db.QueryRow(`SELECT id, assignee, env, name, platform, status, version, context FROM infra WHERE context = ?`, context)
	i := model.Infra{}
	if err := r.Scan(&i.ID, &i.Assignee, &i.Env, &i.Name, &i.Platform, &i.Status, &i.Version, &i.Platform.K8sCluster.Context); err != nil {
		return nil, err
	}

	return &i, nil
}

func (s *SQLiteStore) ListServices() (*[]model.Service, error) {
	r, err := s.db.Query(`SELECT id, name, status, version, image, tag, infra.platform AS platform, infra.name AS platform_name, namespace, env, created_at, assignee, git FROM services INNER JOIN infra ON services.platform = infra.id`)
	if err != nil {
		return nil, err
	}

	svc := model.Service{}
	svcs := []model.Service{}

	for r.Next() {
		if err := r.Scan(&svc.ID, &svc.Name, &svc.Status, &svc.Version, &svc.Image, &svc.Tag, &svc.PlatformKind, &svc.PlatformName, &svc.Namespace, &svc.Env, &svc.CreatedAt, &svc.Assignee, &svc.Git); err != nil {
			return nil, err
		}
		svcs = append(svcs, svc)
	}

	return &svcs, nil
}

func (s *SQLiteStore) GetMetrics(serviceID string) (*[]model.Metric, error) {
	rs, err := s.db.Query(`SELECT id, service_name, metric_name, source, query, active FROM metrics WHERE service_nanme = ?`, serviceID)
	if err != nil {
		return nil, err
	}

	ms := []model.Metric{}
	m := model.Metric{}

	for rs.Next() {
		rs.Scan(&m.ID, &m.ServiceName, &m.MetricName, &m.Source, &m.Query, &m.Active)
		ms = append(ms, m)
	}
	return &ms, nil
}

func (s *SQLiteStore) CreateInfra(infra *model.Infra) error {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec(`INSERT INTO infra (id, assignee, env, name, platform, status, version, context) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, hex.EncodeToString(b), infra.Assignee, infra.Env, infra.Name, infra.Platform.Kind, "notready", infra.Version, infra.Platform.K8sCluster.Context); err != nil {
		fmt.Println("err:", err)
		return err
	}

	return nil
}

func (s *SQLiteStore) ListInfra() (*[]model.Infra, error) {
	rs, err := s.db.Query(`SELECT id, assignee, env, name, platform, status, version, context FROM infra`)
	if err != nil {
		return nil, err
	}

	ms := []model.Infra{}

	for rs.Next() {
		m := model.Infra{}
		if err := rs.Scan(&m.ID, &m.Assignee, &m.Env, &m.Name, &m.Platform.Kind, &m.Status, &m.Version, &m.Platform.K8sCluster.Context); err != nil {
			fmt.Println(err)
			return nil, err
		}
		ms = append(ms, m)
	}

	return &ms, nil
}

func (s *SQLiteStore) GetContextFromServiceID(serviceID string) (string, error) {
	context := ""
	r := s.db.QueryRow(`SELECT context WHERE id = ?`, serviceID)

	if err := r.Scan(&context); err != nil {
		return "", err
	}

	return context, nil
}

func (s *SQLiteStore) CreateEnv(env *model.System) error {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec(`INSERT INTO envs (id, name) VALUES (?, ?)`, hex.EncodeToString(b), env.Name); err != nil {
		fmt.Println("err:", err)
		return err
	}

	return nil
}

func (s *SQLiteStore) CreateTeam(env *model.System) error {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec(`INSERT INTO teams (id, name) VALUES (?, ?)`, hex.EncodeToString(b), env.Name); err != nil {
		fmt.Println("err:", err)
		return err
	}

	return nil
}

func (s *SQLiteStore) CreatePermission(permission *model.System) error {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec(`INSERT INTO permissions (id, name, json, description) VALUES (?, ?, ?, ?)`, hex.EncodeToString(b), permission.Name, permission.JSON, permission.Description); err != nil {
		fmt.Println("err:", err)
		return err
	}

	return nil
}

func (s *SQLiteStore) CreateRole(role *model.System) error {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec(`INSERT INTO roles (id, name, description) VALUES (?, ?, ?)`, hex.EncodeToString(b), role.Name, role.Description); err != nil {
		fmt.Println("err:", err)
		return err
	}

	return nil
}

func (s *SQLiteStore) ListEnvs() (*[]model.System, error) {
	rs, err := s.db.Query(`SELECT id, name FROM envs`)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	envs := []model.System{}

	for rs.Next() {
		env := model.System{}
		if err := rs.Scan(&env.ID, &env.Name); err != nil {
			return nil, err
		}
		envs = append(envs, env)
	}

	return &envs, nil
}

func (s *SQLiteStore) ListTeams() (*[]model.System, error) {
	rs, err := s.db.Query(`SELECT id, name, leader_id FROM teams`)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	teams := []model.System{}

	for rs.Next() {
		team := model.System{}
		if err := rs.Scan(&team.ID, &team.Name, &team.LeaderID); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	return &teams, nil
}

func (s *SQLiteStore) ListPermissions() (*[]model.System, error) {
	rs, err := s.db.Query(`SELECT id, name, json, description FROM permissions`)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	permissions := []model.System{}

	for rs.Next() {
		p := model.System{}
		if err := rs.Scan(&p.ID, &p.Name, &p.JSON, &p.Description); err != nil {
			return nil, err
		}
		permissions = append(permissions, p)
	}

	return &permissions, nil
}

func (s *SQLiteStore) ListRoles() (*[]model.System, error) {
	rs, err := s.db.Query(`SELECT roles.id, roles.name, roles.description, COUNT(role_permissions.permission_id) FROM roles LEFT JOIN role_permissions ON roles.id = role_permissions.role_id GROUP BY roles.id, roles.name, roles.description`)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	roles := []model.System{}

	for rs.Next() {
		role := model.System{}
		if err := rs.Scan(&role.ID, &role.Name, &role.Description, &role.Count); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	for i := range roles {
		ps := []model.Permission{}
		roles[i].Permissions = &ps
		rss, err := s.db.Query(`SELECT permission_id FROM role_permissions WHERE role_id = ?`, roles[i].ID)
		if err != nil {
			fmt.Println("err:", err)
			return nil, err
		}
		for rss.Next() {
			p := model.Permission{}
			rss.Scan(&p.ID)
			ps = append(ps, p)
		}
	}

	return &roles, nil
}

func (s *SQLiteStore) UpdateRole(role *model.System) error {
	if _, err := s.db.Exec(`UPDATE roles SET description = ? WHERE id = ?`, role.Description, role.ID); err != nil {
		fmt.Println("err:", err)
		return err
	}

	if _, err := s.db.Exec(`DELETE FROM role_permissions WHERE role_id = ?`, role.ID); err != nil {
		fmt.Println("err:", err)
		return err
	}

	for _, permission := range *role.Permissions {
		if _, err := s.db.Exec(`INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?)`, role.ID, permission.ID); err != nil {
			fmt.Println("err:", err)
			return err
		}
	}
	return nil
}

func (s *SQLiteStore) ListPermissionsByRoleID(roleID string) (*[]model.System, error) {
	rs, err := s.db.Query(`SELECT permissions.id, permissions.name, permissions.json, permissions.description FROM permissions INNER JOIN role_permissions ON role_permissions.permission_id = permissions.id WHERE role_permissions.role_id = 'roleID'`)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	permissions := []model.System{}

	for rs.Next() {
		p := model.System{}
		if err := rs.Scan(&p.ID, &p.Name, &p.JSON, &p.Description); err != nil {
			return nil, err
		}
		permissions = append(permissions, p)
	}

	return &permissions, nil
}
