package users

import (
	"database/sql"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.db.QueryRow("SELECT id, email, password, created_at FROM users WHERE email=$1", email).
		Scan(&user.Id, &user.Email, &user.Password, &user.Created)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repo) CreateUser(email, password string) error {
	_, err := r.db.Exec("INSERT INTO users (email, password, created_at) VALUES ($1,$2,$3)", email, password, time.Now())
	if err != nil {
		return err
	}
	return nil
}
