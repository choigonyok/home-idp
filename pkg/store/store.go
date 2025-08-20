package store

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

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
