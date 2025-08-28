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

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS services (id TEXT PRIMARY KEY, name TEXT UNIQUE, status TEXT, version TEXT NOT NULL CHECK (version <> ''), image TEXT, tag TEXT, platform TEXT NOT NULL CHECK (platform <> ''), platform_name TEXT NOT NULL CHECK (platform_name <> ''), namespace TEXT, env TEXT NOT NULL CHECK (env <> ''), created_at  DATETIME DEFAULT (CURRENT_TIMESTAMP), assignee TEXT NOT NULL CHECK (assignee <> ''), git TEXT)`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS infra (id TEXT PRIMARY KEY, name TEXT UNIQUE, status TEXT, version TEXT NOT NULL CHECK (version <> ''), platform TEXT NOT NULL CHECK (platform <> ''), env TEXT NOT NULL CHECK (env <> ''), created_at  DATETIME DEFAULT (CURRENT_TIMESTAMP), assignee TEXT NOT NULL CHECK (assignee <> ''), git TEXT)`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS metrics (id TEXT PRIMARY KEY, service_name TEXT UNIQUE, metric_name TEXT, source TEXT, query TEXT, active TEXT)`); err != nil {
		log.Fatalf("failed create table: %v", err)
	}

	return &SQLiteStore{db: db}
}

func (s *SQLiteStore) CreateUser(username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	id := fmt.Sprintf("user_%d", timeNowUnix())
	_, err = s.db.Exec(`INSERT INTO users (id, username, password) VALUES (?, ?, ?)`, id, username, string(hash))
	return err
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

	if _, err := s.db.Exec(`INSERT INTO services (id, name, status, version, image, tag, platform, platform_name, namespace, env, created_at, assignee, git) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, hex.EncodeToString(b), svc.Name, svc.Status, svc.Version, svc.Image, svc.Tag, svc.Platform, svc.PlatformName, svc.Namespace, svc.Env, svc.CreatedAt, svc.Assignee, svc.Git); err != nil {
		return err
	}

	for _, m := range svc.Metrics {
		if _, err := s.db.Exec(`INSERT INTO services (id, service_name, metric_name, source, query, active) VALUES (?, ?, ?, ?, ?)`, m.ID, m.ServiceName, m.MetricName, m.Source, m.Query, m.Active); err != nil {
			return err
		}
	}

	return nil
}

func (s *SQLiteStore) ListServices() (*[]model.Service, error) {
	r, err := s.db.Query(`SELECT id, name, status, version, image, tag, platform, platform_name, namespace, env, created_at, assignee, git FROM services`)
	if err != nil {
		return nil, err
	}

	svc := model.Service{}
	svcs := []model.Service{}

	for r.Next() {
		if err := r.Scan(&svc.ID, &svc.Name, &svc.Status, &svc.Version, &svc.Image, &svc.Tag, &svc.Platform, &svc.PlatformName, &svc.Namespace, &svc.Env, &svc.CreatedAt, &svc.Assignee, &svc.Git); err != nil {
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
	fmt.Println("infra:", infra)
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec(`INSERT INTO infra (id, assignee, env, name, platform, status, version) VALUES (?, ?, ?, ?, ?, ?, ?)`, hex.EncodeToString(b), infra.Assignee, infra.Env, infra.Name, infra.Platform.Kind, "notready", infra.Version); err != nil {
		fmt.Println("err:", err)
		return err
	}

	return nil
}

func (s *SQLiteStore) ListInfra() (*[]model.Infra, error) {
	rs, err := s.db.Query(`SELECT id, assignee, env, name, platform, status, version FROM infra`)
	if err != nil {
		return nil, err
	}

	ms := []model.Infra{}

	for rs.Next() {
		m := model.Infra{}
		if err := rs.Scan(&m.ID, &m.Assignee, &m.Env, &m.Name, &m.Platform.Kind, &m.Status, &m.Version); err != nil {
			fmt.Println(err)
			return nil, err
		}
		ms = append(ms, m)
	}

	fmt.Println(ms)
	return &ms, nil

}
