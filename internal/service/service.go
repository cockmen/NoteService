package service

import (
	"database/sql"
	"notes/internal/notes"
	"notes/internal/users"

	"github.com/labstack/echo/v4"
)

const (
	InvalidParams       = "invalid params"
	InternalServerError = "internal server error"
	InvalidRequest      = "invalid request"
	UserExist           = "user already exist"
	ErrWithPass         = "failed to crypt password"
	UserNotFound        = "user not found"
	InvalidPassword     = "invalid password"
	MissingToken        = "missing token"
	InvalidToken        = "invalid token"
	CantReadQOTD        = "can`t read QOTD"
	ErrWithUnmrshQOTD   = "can`t unmarshal QOTD"
	Unauthorized        = "unauthorized"
)

type Service struct {
	db        *sql.DB
	logger    echo.Logger
	usersRepo *users.Repo
	notesRepo *notes.Repo
}

func NewService(db *sql.DB, logger echo.Logger) *Service {
	svc := &Service{
		db:     db,
		logger: logger,
	}
	svc.initRepositories(db)

	return svc
}

func (s *Service) initRepositories(db *sql.DB) {
	s.usersRepo = users.NewRepo(db)
	s.notesRepo = notes.NewRepo(db)
}

type Response struct {
	Object       any    `json:"object,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

func (r *Response) Error() string {
	return r.ErrorMessage
}

func (s *Service) NewError(err string) (int, *Response) {
	return 400, &Response{ErrorMessage: err}
}
